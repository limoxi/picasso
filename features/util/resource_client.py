# coding: utf8

import json
import sys
import traceback
from time import time
import requests

from features import settings

DEFAULT_TIMEOUT = 30
DEFAULT_GATEWAY_HOST = settings.API_GATEWAY

def unicode_full_stack():
	exc = sys.exc_info()[0]
	stack = traceback.extract_stack()[:-1]  # last one would be full_stack()
	if exc is not None:  # i.e. if an exception is present
		del stack[-1]       # remove call of full_stack, the printed exception
		# will contain the caught exception caller instead
	trc = 'Traceback (most recent call last, REVERSED CALL ORDER):\n'
	stack_str = trc + ''.join(reversed(traceback.format_list(stack)))
	if exc is not None:
		stack_str += '  ' + traceback.format_exc().lstrip(trc)

	return stack_str

class Resource(object):

	@classmethod
	def use(cls, service, gateway_host=DEFAULT_GATEWAY_HOST):
		return cls(service, gateway_host)

	def __init__(self, service='', gateway_host=''):
		self.service = service
		self.gateway_host = gateway_host
		self.__json_data = None
		self.__resource = ''
		if gateway_host.find('://') < 0:
			# 如果没有scheme，则自动补全
			self.gateway_host = "http://%s" % gateway_host

		self.headers = {}
		self.__resp = None

	def set_token(self, token):
		if token == '':
			del self.headers['AUTHORIZATION']
			return
		self.headers.update({
			'AUTHORIZATION': token
		})

	def get(self, options):
		return self.__request(options['resource'], options['data'], 'get')

	def put(self, options):
		return self.__request(options['resource'], options['data'], 'put')

	def post(self, options):
		return self.__request(options['resource'], options['data'], 'post')

	def delete(self, options):
		return self.__request(options['resource'], options['data'], 'delete')

	def __request(self, resource, params, method):
		# 构造url
		"""
		@return is_success,code,data
		"""
		params = params or {}
		self.__resource = resource
		host = self.gateway_host

		resource_path = resource.replace('.', '/')

		service_name = self.service
		self.__target_service = service_name
		if service_name:
			url = '%s/%s/%s/' % (host, service_name, resource_path)
		else:
			# 如果resouce为None，则URL中省略resource。方便本地调试。
			url = '%s/%s/' % (host, resource_path)

		start = time()
		try:
			if method == 'get':
				resp = requests.get(url, params=params, timeout=DEFAULT_TIMEOUT, headers=self.headers)
			else:
				fn = getattr(requests, method)
				self.headers.update({
					'Content-Type': 'application/json;charset=utf-8'
				})
				resp = fn(url, data=json.dumps(params), timeout=DEFAULT_TIMEOUT, headers=self.headers)
			self.__resp = resp

			# 解析响应
			if resp.status_code == 200:
				json_data = json.loads(resp.text)
				self.__json_data = json_data
				code = json_data['code']
				self.__log(code == 200, url, params, method)
				return json_data
			else:
				self.__log(False, url, params, method, 'NetworkError', 'HTTP_STATUS_CODE:' + str(resp.status_code))
				return None

		except requests.exceptions.RequestException as e:
			self.__log(False, url, params, method, str(type(e)), unicode_full_stack())
			return None
		except BaseException as e:
			self.__log(False, url, params, method, str(type(e)), unicode_full_stack())
			return None
		finally:
			stop = time()
			duration = stop - start
			print('expend time {}'.format(duration))

	def __log(self, is_success, url, params, method, failure_type='', failure_msg=''):
		msg = {
			'url': url,
			'params': params,
			'method': method,
			'resource': self.__resource,
			'target_service': self.__target_service,
			'failure_type': failure_type,
			'failure_msg': failure_msg,
		}

		resp = self.__resp

		if resp:
			msg['http_code'] = resp.status_code
			if method == 'get' and is_success:
				msg['resp_text'] = 'stop_record'
			else:
				msg['resp_text'] = self.__json_data
		else:
			msg['http_code'] = ''
			msg['resp_text'] = ''

		print(msg)
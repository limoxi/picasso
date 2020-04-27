# -*- coding: utf-8 -*-

import json
import logging
from jsondiff import diff

class Obj(object):
	def __init__(self):
		pass

def assert_json(expected, actual):
	result = diff(actual, expected)
	for _, item in result.items():
		for k, v in item.items():
			if str(k) not in ['$delete']:
				print('************ASSERT ERROR************\n')
				print str(k), v
				print('************ASSERT ERROR************\n')
				raise RuntimeError(result)

def assert_api_call(response, context):
	if context.text:
		input_data = json.loads(context.text)
		if isinstance(input_data, dict) and 'error' in input_data:
			assert_api_call_failed(response, input_data['error'])
			return False
		elif isinstance(input_data, list) and 'error' in input_data[0]:
			assert_api_call_failed(response, input_data[0]['error'])
			return False
		else:
			assert_api_call_success(response)
			return True
	else:
		assert_api_call_success(response)
		return True


def assert_api_call_success(response):
	"""
	验证api调用成功
	"""
	if 200 != response['code']:
		buf = []
		buf.append('>>>>>>>>>>>>>>> response <<<<<<<<<<<<<<<')
		buf.append(str(response))
		logging.error("API calling failure: %s" % '\n'.join(buf))
	assert 200 == response['code'], "code != 200, call api FAILED!!!!"

def assert_api_call_failed(response, expected_err_code=None):
	"""
	验证api调用失败
	"""
	if 200 == response['code']:
		buf = []
		buf.append('>>>>>>>>>>>>>>> response <<<<<<<<<<<<<<<')
		buf.append(str(response))
		logging.error("API calling not expected: %s" % '\n'.join(buf))
	assert 200 != response['code'], "code == 200, call api NOT EXPECTED!!!!"
	if expected_err_code:
		actual_err_code = str(response.data['errCode'])
		assert expected_err_code in actual_err_code, "errCode != '%s', error code FAILED!!!" % expected_err_code
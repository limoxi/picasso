# -*- coding: utf-8 -*-

import json
import logging

def diff(expected, actual, err_msg=None):
	err_msg = err_msg or []
	if isinstance(expected, dict):
		for k, v in expected:
			if not actual.has_key(k):
				err_msg.append("keyword missing: e.{}".format(k))
			if v != actual[k]:
				err_msg.append("invalid value: e.{}={} => a.{}={}", k, v, k, actual[k])

	elif isinstance(expected, list):
		if len(expected) > len(actual):
			err_msg.append("invalid list length: e.length={} => a.length={}".format(len(expected), len(actual)))
		else:
			for i, ee in enumerate(expected):
				err_msg += diff(ee, actual[i], err_msg)
	else:
		if expected != actual:
			err_msg.append("invalid value: e={} => a={}".format(expected, actual))

	return err_msg

class Obj(object):
	def __init__(self):
		pass

def assert_json(expected, actual):
	err_msg = diff(expected, actual)
	if len(err_msg) > 0:
		print "expected: ", expected
		print "actual: ", actual
		raise Exception('assert error >>>>>>>>> \n' + '\n'.join(err_msg))

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
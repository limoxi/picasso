# -*- coding: utf-8 -*-

def print_request(path, method, payload):
	try:
		print("="*20 + "Start print BDD Request" + "="*20)
		print("Request:")
		print("    Method: %s" %  method)
		print("    Url: %s" % path)
		print("    Params:")
		for key, value in payload.items():
			print("        %s: %s," % (key, value))
	except Exception as e:
		print("print request error: %s" % e)
	finally:
		print("="*20 + "Finish print BDD Request" + "="*20)


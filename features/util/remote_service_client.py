# coding: utf8

from features.util.resource_client import Resource
from features import settings

class RemoteServiceClient(object):

	def __init__(self):
		self.token = None
		self.user_phone = ''
		self.resource = Resource.use(settings.SERVICE_NAME)

	def login(self, phone='13111111111'):
		print('login with phone: {}'.format(phone))
		resp = self.resource.put({
			'resource': 'user.logined_user',
			'data': {
				'phone': phone,
				'password': settings.BDD_USER_PWD
			}
		})
		if resp and resp['code'] == 200:
			self.token = resp['data']['token']
			self.user_phone = phone
			self.resource.set_token(self.token)
		else:
			assert False
		return self

	def logout(self):
		print('logout current user {}'.format(self.user_phone))
		self.token = None
		self.user_phone = ''
		self.resource.set_token('')
		return self

	def get(self, resource, data=None, jwt_token=None):
		return self.resource.get({
			'resource': resource,
			'data': data
		})

	def put(self, resource, data=None, jwt_token=None):
		return self.resource.put({
			'resource': resource,
			'data': data
		})

	def post(self, resource, data=None, jwt_token=None):
		return self.resource.post({
			'resource': resource,
			'data': data
		})

	def delete(self, resource, data=None, jwt_token=None):
		return self.resource.delete({
			'resource': resource,
			'data': data
		})
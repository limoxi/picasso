# coding: utf8
from features import settings

from features.util import bdd_util
from features.util.remote_service_client import RemoteServiceClient

USERS = {
	"manager": "13111111111",
	"zhang3": "13111111112",
	"li4": "13111111113",
	"wang5": "13111111114",
	"zhao6": "13111111115"
}

remote_client = RemoteServiceClient()

class Obj(object):
	def __init__(self):
		pass

def create_users():
	"""
	创建所有bdd用户
	"""
	init_users()


def init_users():
	"""
	创建普通user
	"""
	for name, phone in USERS.items():
		data = {
			"phone": phone,
			"passsword": settings.BDD_USER_PWD,
		}
		resp = remote_client.put('user.registered_user', data)
		bdd_util.assert_api_call_success(resp)
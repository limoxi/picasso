# coding: utf8
from copy import copy

from features import settings

from features.util import bdd_util
from features.util.remote_service_client import RemoteServiceClient

from features.util.db_util import SQLService

USERS = {
	"manager": "13111111111",
	"zhang3": "13111111112",
	"li4": "13111111113",
	"wang5": "13111111114",
	"zhao6": "13111111115"
}

USERNAME2PHONE = {}

remote_client = RemoteServiceClient()

def init_username2phone():
	global USERNAME2PHONE
	USERNAME2PHONE = copy(USERS)

class Obj(object):
	def __init__(self):
		pass

def init_bdd_users():
	"""
	初始化预置user
	"""
	sql = """
		delete from auth_user;
	"""
	SQLService.use().execute_sql(sql)
	for name, phone in USERS.items():
		data = {
			"phone": phone,
			"password": settings.BDD_USER_PWD,
		}
		resp = remote_client.put('user.registered_user', data)
		bdd_util.assert_api_call_success(resp)

def get_user_id(username):
	phone = USERNAME2PHONE[username]
	if phone is None:
		raise Exception("invalid username: " + username)
	sql = """
		select id from auth_user where phone='{}';
	""".format(phone)
	record = SQLService.use().execute_sql(sql).fetchone()
	return record[0]
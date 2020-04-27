# coding: utf8
import json

from behave import *

from features import settings
from features.util import user_util, bdd_util


@when(u"'{username}'注册用户")
def step_impl(context, username):
	data = json.loads(context.text)
	response = context.client.put('user.registered_user', {
		'phone': data['phone'],
		'password': data.get('password', settings.BDD_USER_PWD),
	})
	bdd_util.assert_api_call_success(response)
	user_util.USERNAME2PHONE[username] = data['phone']

@then(u"'{username}'可以使用密码'{pwd}'登陆系统")
def step_impl(context, username, pwd):
	response = context.client.put('user.logined_user', {
		'phone': USERNAME2PHONE[username],
		'password': pwd
	})
	bdd_util.assert_api_call_success(response)
	resp_user = response['data']
	if resp_user['token'] == "":
		assert False

@given(u"{username}登录系统")
def step_impl(context, username):
	response = context.client.put('user.logined_user', {
		'phone': user_util.USERNAME2PHONE[username],
		'password': settings.BDD_USER_PWD
	})
	bdd_util.assert_api_call_success(response)
	resp_user = response['data']
	if resp_user['token'] == "":
		assert False
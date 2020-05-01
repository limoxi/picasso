# coding: utf8

import json
from behave import *

from features import settings
from features.util import user_util, bdd_util
from features.util.db_util import SQLService


def get_space_id(space_name, username=None):
	sql = u"""
		select id from space_space where name='{}' 
	""".format(space_name)
	if username is not None:
		user_id = user_util.get_user_id(username)
		sql += " and user_id={} ".format(user_id)
	record = SQLService.use().execute_sql(sql).fetchone()
	return record[0]

def update_space_code(space_id, code):
	sql = """
		update space_space set code='{}' 
		where id={};
	""".format(code, space_id)
	SQLService.use().execute_sql(sql)

@when(u"'{username}'创建空间'{space_name}'")
def step_impl(context, username, space_name):
	response = context.client.put('space.space', {
		'name': space_name
	})
	bdd_util.assert_api_call_success(response)

@then(u"'{username}'可以查看自己的空间列表")
def step_impl(context, username):
	expected = json.loads(context.text)
	response = context.client.get('space.spaces', {})
	actual = response['data']['spaces']
	bdd_util.assert_json(expected, actual)

@when(u"'{username}'为空间'{space_name}'生成邀请码'{code}'")
def step_impl(context, username, space_name, code):
	space_id = get_space_id(space_name, username)
	response = context.client.put('space.code', {
		"space_id": space_id,
	})
	bdd_util.assert_api_call_success(response)
	update_space_code(space_id, code)

@when(u"'{member_name}'使用邀请码'{code}'加入空间'{space_name}'")
def step_impl(context, member_name, code, space_name):
	response = context.client.put('space.member', {
		"space_id": get_space_id(space_name),
		"code": code
	})
	bdd_util.assert_api_call_success(response)

@then(u"'{username}'可以查看空间'{space_name}'的成员列表")
def step_impl(context, username, space_name):
	expected = json.loads(context.text)
	response = context.client.get('space.members', {
		"space_id": get_space_id(space_name, username)
	})
	actual = response['data']['members']
	for member in actual:
		if member['nick_name'] == '':
			sql = """
				select phone from auth_user where id={}
			""".format(member['user_id'])
			record = SQLService.use().execute_sql(sql).fetchone()
			for nick_name, phone in user_util.USERNAME2PHONE.items():
				if phone == record[0]:
					member['nick_name'] = nick_name
					break
	bdd_util.assert_json(expected, actual)
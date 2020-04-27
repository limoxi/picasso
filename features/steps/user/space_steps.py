# coding: utf8

import json

from behave import *

from features import settings
from features.util import user_util, bdd_util

from features.util.db_util import SQLService


def get_space_id(username, space_name):
	user_id = user_util.get_user_id(username)
	sql = """
		select * from space_space where user_id={} and name='{}'; 
	""".format(user_id, space_name)
	record = SQLService.use().execute_sql(sql).fetchone()
	return record[0]

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

@when(u"'{username}'邀请'{member_name}'成为空间'{space_name}'的成员")
def step_impl(context, username, member_name, space_name):
	response = context.client.put('space.invited_member', {
		"space_id": get_space_id(username, space_name),
		"user_id": user_util.get_user_id(member_name)
	})
	bdd_util.assert_api_call_success(response)

@when(u"'{member_name}'同意成为'{username}'的空间'{space_name}'的成员")
def step_impl(context, member_name, username, space_name):
	response = context.client.put('space.passed_inviting', {
		"space_id": get_space_id(username, space_name)
	})
	bdd_util.assert_api_call_success(response)

@when(u"'{member_name}'拒绝成为'{username}'的空间'{space_name}'的成员")
def step_impl(context, member_name, username, space_name):
	response = context.client.put('space.rejected_inviting', {
		"space_id": get_space_id(username, space_name)
	})
	bdd_util.assert_api_call_success(response)

@then(u"'{username}'可以查看空间'{space_name}'的成员列表")
def step_impl(context, username, space_name):
	expected = json.loads(context.text)
	response = context.client.get('space.members', {
		"space_id": get_space_id(username, space_name)
	})
	actual = response['data']['members']
	bdd_util.assert_list(expected, actual)
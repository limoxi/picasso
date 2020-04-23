# coding: utf8
import os
import sys

from features.util.db_util import SQLService
from features.util.remote_service_client import RemoteServiceClient
from features.util import bdd_util, user_util
from features import settings

path = os.path.abspath(os.path.join('.', '..'))
sys.path.insert(0, path)

if settings.DATABASES['HOST'] not in ['db.dev.com', '127.0.0.1']:
	raise RuntimeError("run BDD when connect local database")

def __clear_data():
	"""
	清空应用数据
	"""
	sqls = ["""
		delete from user_user;
	"""]
	teamdo_db = SQLService.use()
	for sql in sqls:
		teamdo_db.execute_sql(sql)

def before_all(context):

	pass

def after_all(context):
	pass


def before_scenario(context, scenario):
	context.scenario = scenario
	context.client = RemoteServiceClient()
	__clear_data()

def after_scenario(context, scenario):
	if hasattr(context, 'client') and context.client and context.client.user:
		context.client.logout()
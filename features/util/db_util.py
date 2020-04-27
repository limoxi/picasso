# coding: utf8
from playhouse.db_url import connect
from features import settings

DB_DICT = dict()

def get_db(name):
	global DB_DICT
	if DB_DICT.get(name):
		return DB_DICT[name]
	db_config = settings.DATABASES
	url = '{engine}://{user}:{pwd}@{host}:{port}/{db_name}'.format(
		engine=db_config['ENGINE'],
		user=db_config['USER'],
		pwd=db_config['PASSWORD'],
		host=db_config['HOST'],
		port=db_config['PORT'],
		db_name=name
	)
	print('bdd connect db: %s' % url)
	db = connect(url)
	db.connect()
	DB_DICT[name] = db
	return db

class SQLService(object):
	"""
	直接执行sql的服务
	"""
	__slots__ = (
		'__db',  # database实例
	)

	def __init__(self, db):
		self.__db = db

	@classmethod
	def use(cls, db_name='picasso'):
		db = get_db(db_name)
		if not db:
			raise Exception
		return cls(db)

	def execute_sql(self, sql, for_test=False):
		cursor = None
		for s in sql.split(';'):
			s = s.strip()
			if not s:
				continue
			print (u'execute sql: {}'.format(s))
			if not for_test:
				cursor = self.__db.execute_sql(s)
		return cursor

	def set_foreign_key_checks(self, value=True):
		sql = """
			set foreign_key_checks={};
		""".format(int(value))
		self.execute_sql(sql)
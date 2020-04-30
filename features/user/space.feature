# Created by limoxi at 2020/4/25 0025
Feature: 用户空间管理
	Background:

	@picasso @space
	Scenario: 1、用户可以创建空间
		Given zhang3登录系统
		Then 'zhang3'可以查看自己的空间列表
		"""
		[]
		"""
		When 'zhang3'创建空间'张小三'
		Then 'zhang3'可以查看自己的空间列表
		"""
		[{
			"name": "张小三"
		}]
		"""
		Given li4登录系统
		When 'li4'创建空间'李小四'
		And 'li4'创建空间'李小五'
		Then 'li4'可以查看自己的空间列表
		"""
		[{
			"name": "李小五"
		}, {
			"name": "李小四"
		}]
		"""

	@picasso @space @1
	Scenario: 2、空间管理员可以邀请其他用户成为空间一员
		Given zhang3登录系统
		When 'zhang3'创建空间'张小三'
		And 'zhang3'创建空间'张小四'
		And 'zhang3'为空间'张小三'生成邀请码'xxx1'
		And 'zhang3'为空间'张小四'生成邀请码'xxx2'
		Given li4登录系统
		When 'li4'使用邀请码'xxx1'加入空间'张小三'
		Then 'li4'可以查看自己的空间列表
		"""
		[{
			"name": "张小三"
		}]
		"""
		Given zhao6登录系统
		When 'zhao6'使用邀请码'xxx2'加入空间'张小四'
		Then 'zhao6'可以查看自己的空间列表
		"""
		[{
			"name": "张小四"
		}]
		"""
		Given zhang3登录系统
		Then 'zhang3'可以查看空间'张小三'的成员列表
		"""
		[{
			"username": "li4",
			"status": "member"
		}, {
			"username": "zhang3",
			"status": "manager"
		}]
		"""
		And 'zhang3'可以查看空间'张小四'的成员列表
		"""
		[{
			"username": "zhao6",
			"status": "member"
		}, {
			"username": "zhang3",
			"status": "manager"
		}]
		"""
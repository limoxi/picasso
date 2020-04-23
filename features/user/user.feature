Feature: 用户可以注册成为会员并登陆系统
    
    @picasso @user @user.register
    Scenario: 1、用户使用手机号注册
        When 'xiaohua'注册用户
        """
        {
            "phone": "18011111111",
            "password": "test"
        }
        """
        Then 'xiaohua'可以使用密码'test'登陆系统
# 依赖
```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/viper
...gorm...
go mod tidy
```


# 使用sqlite做数据迁移验证
## sqlite connect 格式
- file:./example.db?cache=shared&mode=rw
# 使用mysql做生产数据库迁移验证
## mysql dsn 格式
- root:123456@tcp(127.0.0.1:3306)/prod?charset=utf8mb4&parseTime=True&loc=Local

# 数据库表说明(后面带v是版本意思)
```excel
patient	              存储患者的基本信息
doctor	              存储医生的基本信息
visit	              存储患者的就诊信息
order	              存储医生为患者开具的医嘱
labtest	              存储检验检查单项目表
medical_record	      存储患者的病历信息
userv2	              存储系统用户（医生、管理员等）的账号信息
admission_notice	  存储患者入院通知单信息
western_medicline	  存储西药信息表
chinese_medicline	  存储中草药信息表
ICD_10	              存储诊断编码表
```


# 来自数据库的字段todo
```bash
首页展示
{
首诊时间
接诊时间
患者姓名
门诊病历号
患者性别
患者年龄
患者职业
患者联系地址
患者联系电话
}

诊断展示?
{
诊断路径
}

患者基本信息展示
个人信息:
    姓名
    性别
    出生日期
    证件
    联系电话
    联系地址
    国籍
    民族
    宗教
    学历
    学位
    职业分类
    工作单位名称
    工作单位电话号码
    单位地址（选址combox）
    婚姻状况
    籍贯
    出生地（选址combox）
    贫困等级
    联系人姓名

健康摘要:
    过敏史
    粉尘史
    健康状况
    病理状态
    身高
    体重
    BMI
    发热
    血压
    脉搏
    心率
    呼吸
    血糖
    血氧饱和度
    血型
    RH血型
    是否到过疫区
    是否做过新冠检测
    特殊人员类型
    危重级别





```

# ui控制todo
1. 首页GridView展示5条数据，根据时间筛选


# api todo
1. 密码登录和加密存储（简单）
2. 患者历次病历信息（简单）
3. 药品信息 （todo）


# ref link:
- https://gitee.com/EEPPEE_admin/doctor-api-backend
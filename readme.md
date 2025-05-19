> [!WARNING]  
> 全部代码几乎由ai生成，不写了，或者迁移到gitee.com/eeppee_admin/doctor-api-backend

# 思路
1. drop掉所有表，生成数据库model
2. 生成model 的fake data
3. 编写api，含GET POST
4. docker运行起来
4. 测试api调用情况

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



# 内网穿透guide (fixed)
```console
$cpolar 8080
Forwarding          http://2ece470b.r7.cpolar.top -> http://localhost:8080
# curl测试
curl http://2ece470b.r7.cpolar.top/api/medicines
```

# github first push guide (fixed)
```bash
# met error:
#remote: Resolving deltas: 100% (10/10), done.
# remote: fatal: did not receive expected object bcfa1353e3b40252581153fe39fc58a1cb8518df
# error: remote unpack failed: index-pack failed
# To https://github.com/eeppee-admin/-api-.git
#  ! [remote rejected] master -> master (failed)
# error: failed to push some refs to 'https://github.com/eeppee-admin/-api-.git'
# do:
git fetch --unshallow 
# then github is my new remote
git push github master
```

# ref link:
- https://gitee.com/EEPPEE_admin/doctor-api-backend
- https://www.cpolar.com/
- github action学习:https://shipengqi.github.io/golang-learn/docs/project/09_actions/#:~:text=GitHub%20Actions%20%E6%98%AF%20GitHub%20%E4%B8%BA%E6%89%98%E7%AE%A1%E5%9C%A8%20github.com%20%E7%AB%99%E7%82%B9%E7%9A%84%E9%A1%B9%E7%9B%AE%E6%8F%90%E4%BE%9B%E7%9A%84%E6%8C%81%E7%BB%AD%E9%9B%86%E6%88%90%E6%9C%8D%E5%8A%A1%E3%80%82%20%E5%9C%A8%E6%9E%84%E5%BB%BA%E6%8C%81%E7%BB%AD%E9%9B%86%E6%88%90%E4%BB%BB%E5%8A%A1%E6%97%B6%EF%BC%8C%E9%9C%80%E8%A6%81%E5%AE%8C%E6%88%90%E5%BE%88%E5%A4%9A%E6%93%8D%E4%BD%9C%EF%BC%8C%E6%AF%94%E5%A6%82%E5%85%8B%E9%9A%86%E4%BB%A3%E7%A0%81%E3%80%81%E7%BC%96%E8%AF%91%E4%BB%A3%E7%A0%81%E3%80%81%E8%BF%90%E8%A1%8C%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95%E3%80%81%E6%9E%84%E5%BB%BA%E5%92%8C%E5%8F%91%E5%B8%83%E9%95%9C%E5%83%8F%E7%AD%89%E3%80%82,Actions%20%E6%98%AF%E5%8F%AF%E4%BB%A5%E5%85%B1%E4%BA%AB%E7%9A%84%EF%BC%8C%E5%BC%80%E5%8F%91%E8%80%85%E5%8F%AF%E4%BB%A5%E5%B0%86%20Actions%20%E4%B8%8A%E4%BC%A0%E5%88%B0%20GitHub%20%E7%9A%84%20Actions%20%E5%B8%82%E5%9C%BA%E3%80%82

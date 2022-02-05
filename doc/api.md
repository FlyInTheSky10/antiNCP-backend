# 健康上报系统后端 API 文档

## 数据交换格式

### 身份验证

在 Header 中加入 `Authorization` 字段进行验证，将获取的 JWT 令牌作为 Bearer Token 加入该字段的值，例如：

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U
```

以下接口中，标题带有 `*` 标记的为需要身份验证的接口。

### 响应格式

响应使用 JSON 格式，例如：

```json
{
    "success": true,
    "info": '',
    "data": {
        // ...
    }
}
```

### URL 前缀

文档中所有接口 URL 都包含前缀 `/api/v1`。

## 值约定

### time

所有时间使用 Unix 时间戳。

### status

- `0`: 未审核，无信息
- `1`: 已通过
- `2`: 已拒绝

### accepted

- `0`: 未审核，无信息
- `1`: 已通过
- `2`: 已拒绝

### admin

- `0`：非 admin
- `1`：admin

### vaccine_stage

- `0`： 未接种
- `1`：一针
- `2`：两针
- `3`：加强针

### health

- `0`：正常
- `1`：异常

### travel_history

- `0`：正常
- `1`：异常

### health_code

- `0`：正常
- `1`：异常

### travel_code

- `0`：正常
- `1`：异常

## 用户 API

### 获取公钥 GET /user/publicKey?id={ID学号}

为了保证密码安全，登录时用 RSA 加密密码传输，获取一次公钥有效期 15 分钟。公钥格式为 PKSC1，加密应使用 PKCS #1 v1.5

响应：

```json
{
  data: {
    publickey: "base64 编码的 1024 位 RSA 公钥"
  },
  success: true,
  info: ''
}
```

### 注册用户 PUT /user/register

请求：

```json
{
    id: "2021150210",
    "password": "加密后的密码",
    name: "名字",
    academy: "学院",
    idNumber: "证件号",
    phoneNumber: "电话号码"
}
```

响应：successfully add user

### 登录获取令牌 GET /user/token?id={ID学号}&password={加密密码}

获取 JWT 令牌，调用前使用获取公钥接口。

响应：

```json
{
    data: {
        token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjIwMjExMTAyMTMiLCJleHAiOjE2NDIzNjgzNjV9.8fPV3wb-pgrnEQZYxlvBjEFe5vufFPNBTf0wIKpQ90g',
        expire: 1642368365
    },
    success: true,
    info: ''
}
```

### *查看状态 GET /user/status

请求：无

响应：

```json
{
  data: { Status: 3, VerifyId: 2021150213, VerifyTime: 1642323819 },
  success: true,
  info: ''
}
```

### *获取用户 GET /user/

请求：无

响应：

```json
{
  data: { Name:xxx, Id:xxx, Academy:xxx, IdNumber:xxx, PhoneNumber:xxx },
  success: true,
  info: ''
}
```

## 提交 API

### *提交申请 PUT /code/submitCode

请求：

```json
{
    vaccineStage: 3,
    health: 0,
    travelHistory: 0,
    healthCode: 0,
    travelCode: 0
}
```

响应：无

### *用户查看所有申请 GET /code/viewSubmission?startIndex={index}

返回 index [index, index + 10] 从0开始

响应：

```json
{
  data: [
    {
      Id: '2021110211',
      I: 1,
      VaccineStage: 3,
      Health: 1,
      TravelHistory: 1,
      HealthCode: 0,
      TravelCode: 1,
      SubmitTime: 1642236585,
      Accepted: 0
    },
    {
      Id: '2021110211',
      I: 2,
      VaccineStage: 3,
      Health: 1,
      TravelHistory: 0,
      HealthCode: 0,
      TravelCode: 1,
      SubmitTime: 1642236608,
      Accepted: 0
    }
  ],
  success: true,
  info: ''
}
```

### *用户查看所有提交数量 GET /code/viewSubmissionNumber

### *Admin查看所有提交 GET /code/allSubmission?startIndex={index}

(需要是Admin) 返回 index [index, index + 10] 从0开始

响应：

```json
{
  data: [
    {
      Id: '2021110211',
      I: 1,
      VaccineStage: 3,
      Health: 1,
      TravelHistory: 1,
      HealthCode: 0,
      TravelCode: 1,
      SubmitTime: 1642236585,
      Accepted: 0
    },
    {
      Id: '2021110215',
      I: 2,
      VaccineStage: 3,
      Health: 1,
      TravelHistory: 0,
      HealthCode: 0,
      TravelCode: 1,
      SubmitTime: 1642236608,
      Accepted: 0
    }
  ],
  success: true,
  info: ''
}
```

### *Admin查看所有提交数量 GET /code/allSubmissionNumber

### *Admin审核提交 PUT /code/verifySubmission

(需要是Admin) 请求：

```json
{
    i: 3, // 第 {i} 个提交
    status: 1 // 使其状态改变
}
```

响应：无

## Admin API

### *Admin查看所有注册用户状态 GET /admin/allStatus?startIndex={index}

### *Admin查询所有注册用户数量 GET /admin/allUserNumber


// It's a full demo for pica.

// Init vars
name = 'demo'
description = 'This is a demo for pica.'
version = '0.0.1'
author = '<jeremaihloo@gmail.com> jeremaihloo'
baseUrl = 'http://localhost:8080'

headers = {
    'Content-Type' = 'application/json'
}

// Apis format: [method] [path] [description]

// GET /api/users 获取用户列表
headers['Authorization'] = 'slfjaslkfjlasdjfjas=='

// POST /api/users 新建用户
post = {
    'name' = 'test'
    'age' = 10
}
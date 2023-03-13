import {request, requestBody} from "./http";

// 登录
export const login = (
    username: string,
    password: string,
    captcha: string,
) => request('/api/session', 'POST', null, {
    username: username,
    password: password,
    captcha: captcha,
});
// 注销
export const logout = () => request('/api/session', 'DELETE');

export const queryDashboard = () => request('/api/dashboard', 'GET', null, null)

// 查询文章
export const queryArticle = (
    title: string | null,
    categoryId: number | null,
    url: string | null,
    enableComment: boolean | null,
    status: number | null,
    pageNum: number,
    pageSize: number,
    orderName: string | null,
    orderMethod: string | null,
) => request('/api/article', 'GET', {
    title: title,
    categoryId: categoryId,
    url: url,
    enableComment: enableComment,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
    orderName: orderName,
    orderMethod: orderMethod,
}, null)
// 查询文章
export const queryArticleInfo = (id: number) => request('/api/article/' + id, 'GET', null, null)
// 构建文章表单
export const createFormData = (
    title: string,
    categoryId: number | null,
    url: string | null,
    tags: string[],
    content: string,
    weight: number | null,
    enableComment: boolean,
    status: number,
): FormData => {
    const formData = new FormData();
    formData.append("title", title + '')
    formData.append("categoryId", categoryId + '')
    if (url) {
        formData.append("url", url + '')
    }
    if (tags) {
        for (let i = 0; i < tags.length; i++) {
            formData.append("tags", tags[i]);
        }
    }
    formData.append("content", content + '')
    if (weight) {
        formData.append("weight", weight + '')
    }
    formData.append("enableComment", enableComment + '')
    formData.append("status", status + '')
    return formData
}
// 新增文章
export const addArticle = (
    title: string,
    categoryId: number | null,
    url: string | null,
    tags: string[],
    content: string,
    weight: number | null,
    enableComment: boolean,
    status: number,
) => requestBody('/api/article', 'POST', null, createFormData(
    title,
    categoryId,
    url,
    tags,
    content,
    weight,
    enableComment,
    status,
))
// 更新文章
export const updateArticle = (
    id: number,
    title: string,
    categoryId: number | null,
    url: string | null,
    tags: string[],
    content: string,
    weight: number | null,
    enableComment: boolean,
    status: number,
) => requestBody('/api/article/' + id, 'PUT', null, createFormData(
    title,
    categoryId,
    url,
    tags,
    content,
    weight,
    enableComment,
    status,
))
// 删除文章
export const deleteArticle = (id: number) => request('/api/article/' + id, 'DELETE', null, null)

// 查询评论
export const queryComment = (
    article: string | null,
    author: string | null,
    content: string | null,
    email: string | null,
    ip: string | null,
    status: number | null,
    pageNum: number,
    pageSize: number,
    orderName: string | null,
    orderMethod: string | null,
) => request('/api/comment', 'GET', {
    article: article,
    author: author,
    content: content,
    email: email,
    ip: ip,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
    orderName: orderName,
    orderMethod: orderMethod,
}, null)
// 新增评论(回复)
export const addCommentByReplay = (
    articleId: string,
    parentId: string,
    content: string,
    emailNotify: boolean,
) => request('/api/comment', 'POST', null, {
    articleId: articleId,
    parentId: parentId,
    content: content,
    emailNotify: emailNotify,
})
// 更新评论(审核通过)
export const updateCommentToVerify = (
    id: number,
) => request('/api/comment/' + id, 'PUT', null, {
    status: 1,
})
// 删除文章
export const deleteComment = (id: number) => request('/api/comment/' + id, 'DELETE', null, null)

// 查询分类
export const queryCategory = (
    name: string | null,
    pageNum: number,
    pageSize: number,
    orderName: string | null,
    orderMethod: string | null,
) => request('/api/category', 'GET', {
    name: name,
    pageNum: pageNum,
    pageSize: pageSize,
    orderName: orderName,
    orderMethod: orderMethod,
}, null)
// 查询分类选项
export const queryCategoryOptions = () => queryCategory(null, 1, 0x7fffffff, null, null).then((data) => {
    return data.data.list.map((item: any) => {
        return {
            value: item.id,
            label: item.name,
        }
    })
})
// 新增分类
export const addCategory = (
    name: string,
) => request('/api/category', 'POST', null, {
    name: name,
})
// 更新分类
export const updateCategory = (
    id: number,
    name: string,
) => request('/api/category/' + id, 'PUT', null, {
    name: name,
})
// 删除分类
export const deleteCategory = (id: number) => request('/api/category/' + id, 'DELETE', null, null)

// 查询标签
export const queryTag = (
    name: string | null,
    pageNum: number,
    pageSize: number,
    orderName: string | null,
    orderMethod: string | null,
) => request('/api/tag', 'GET', {
    name: name,
    pageNum: pageNum,
    pageSize: pageSize,
    orderName: orderName,
    orderMethod: orderMethod,
}, null)
// 新增标签
export const addTag = (
    name: string,
) => request('/api/tag', 'POST', null, {
    name: name,
})
// 更新标签
export const updateTag = (
    id: number,
    name: string,
) => request('/api/tag/' + id, 'PUT', null, {
    name: name,
})
// 删除标签
export const deleteTag = (id: number) => request('/api/tag/' + id, 'DELETE', null, null)

// 查询文件
export const queryUpload = (
    displayName: string | null,
    type: string | null,
    pageNum: number,
    pageSize: number,
    orderName: string | null,
    orderMethod: string | null,
) => request('/api/upload', 'GET', {
    name: displayName,
    type: type,
    pageNum: pageNum,
    pageSize: pageSize,
    orderName: orderName,
    orderMethod: orderMethod,
}, null)
// 上传文件
export const addUpload = (file: any) => {
    const formData = new FormData();
    formData.append("file", file);
    return requestBody('/api/upload', 'POST', null, formData)
}
// 上传文件
export const updateUpload = (id: number, file: any) => {
    const formData = new FormData();
    formData.append("file", file);
    return requestBody('/api/upload/' + id, 'PUT', null, formData)
}
// 删除文件
export const deleteUpload = (id: number) => request('/api/upload/' + id, 'DELETE', null, null)

// 查询友链
export const queryLink = (
    title: string | null,
    url: string | null,
    status: number | null,
    pageNum: number,
    pageSize: number,
    orderName: string | null,
    orderMethod: string | null,
) => request('/api/link', 'GET', {
    title: title,
    url: url,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
    orderName: orderName,
    orderMethod: orderMethod,
}, null)
// 新增友链
export const addLink = (
    title: string,
    url: string,
    weight: number,
    status: string,
) => request('/api/link', 'POST', null, {
    title: title,
    url: url,
    weight: weight,
    status: status,
})
// 更新友链
export const updateLink = (
    id: number,
    title: string,
    url: string,
    weight: number,
    status: string,
) => request('/api/link/' + id, 'PUT', null, {
    title: title,
    url: url,
    weight: weight,
    status: status,
})
// 删除友链
export const deleteLink = (id: number) => request('/api/link/' + id, 'DELETE', null, null)

// 查询短链
export const queryUrl = (
    key: string | null,
    url: string | null,
    description: string | null,
    status: number | null,
    pageNum: number,
    pageSize: number,
    orderName: string | null,
    orderMethod: string | null,
) => request('/api/url', 'GET', {
    key: key,
    url: url,
    description: description,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
    orderName: orderName,
    orderMethod: orderMethod,
}, null)
// 新增短链
export const addUrl = (
    key: string,
    url: string,
    description: string,
    status: string,
) => request('/api/url', 'POST', null, {
    key: key,
    url: url,
    description: description,
    status: status,
})
// 更新短链
export const updateUrl = (
    id: number,
    key: string,
    url: string,
    description: string,
    status: string,
) => request('/api/url/' + id, 'PUT', null, {
    key: key,
    url: url,
    description: description,
    status: status,
})
// 删除短链
export const deleteUrl = (id: number) => request('/api/url/' + id, 'DELETE', null, null)

// 查询配置
export const queryConfig = (
    key: string | null,
    value: string | null,
    description: string | null,
    pageNum: number,
    pageSize: number,
    orderName: string | null,
    orderMethod: string | null,
) => request('/api/config', 'GET', {
    key: key,
    value: value,
    description: description,
    pageNum: pageNum,
    pageSize: pageSize,
    orderName: orderName,
    orderMethod: orderMethod,
}, null)
// 新增配置
export const addConfig = (
    key: string,
    value: string,
    description: string,
) => request('/api/config', 'POST', null, {
    key: key,
    value: value,
    description: description,
})
// 更新配置
export const updateConfig = (
    key: string,
    value: string,
    description: string,
) => request('/api/config/' + key, 'PUT', null, {
    key: key,
    value: value,
    description: description,
    status: status,
})
// 删除配置
export const deleteConfig = (key: string) => request('/api/config/' + key, 'DELETE', null, null)

// 功能-查询SMTP
export const querySMTP = () => request('/api/features/smtp', 'GET', null, null)
// 功能-更新SMTP
export const updateSMTP = (
    enable: boolean,
    hostname: string | null,
    port: number | null,
    ssl: boolean | null,
    username: string | null,
    password: string | null,
    fromName: string | null,
    fromEmail: string | null,
) => request('/api/features/smtp', 'PUT', null, {
    enable: enable,
    hostname: hostname,
    port: port,
    ssl: ssl,
    username: username,
    password: password,
    fromName: fromName,
    fromEmail: fromEmail,
})
// 功能-查询SMTP
export const testSMTP = (
    email: string,
) => request('/api/features/smtp/test', 'POST', null, {
    email: email,
})


// 查询个人资料
export const queryProfile = () => request('/api/user/profile', 'GET', null, null)
// 修改个人资料
export const updateProfile = (
    avatar: string,
    nickname: string,
    description: string,
    email: string,

    github: string,
    weibo: string,
    google: string,
    twitter: string,
    facebook: string,
    stackOverflow: string,
    youtube: string,
    instagram: string,
    skype: string,
) => request('/api/user/profile', 'PUT', null, {
    avatar: avatar,
    nickname: nickname,
    description: description,
    email: email,

    github: github,
    weibo: weibo,
    google: google,
    twitter: twitter,
    facebook: facebook,
    stackOverflow: stackOverflow,
    youtube: youtube,
    instagram: instagram,
    skype: skype,
})
// 修改密码
export const updatePassword = (
    oldPassword: string,
    newPassword: string,
) => request('/api/user/password', 'PUT', null, {
    oldPassword: oldPassword,
    newPassword: newPassword,
})


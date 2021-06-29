import {request} from "./http";

// 登录
export const login = (username: string, password: string, verifyCode: string) => request('/api/session', 'POST', null, {
    username: username,
    password: password,
    verifyCode: verifyCode,
});
// 注销
export const logout = () => request('/api/session', 'DELETE');

// 查询文章
export const queryArticle = (
    title: string | null,
    category: number | null,
    format: number | null,
    comment: number | null,
    status: number | null,
    pageNum: number,
    pageSize: number,
) => request('/api/article', 'GET', {
    title: title,
    category: category,
    format: format,
    comment: comment,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
}, null)
// 新增文章
export const addArticle = (title: string) => request('/api/article', 'POST', null, {
    title: title,
})
// 更新文章
export const updateArticle = (id: number, title: string) => request('/api/article/' + id, 'PUT', null, {
    title: title,
})
// 更新文章状态
export const updateArticleStatus = (id: number, status: number) => request('/api/article/' + id + '/status', 'PUT', null, {
    status: status,
})
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
) => request('/api/comment', 'GET', {
    article: article,
    author: author,
    content: content,
    email: email,
    ip: ip,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
}, null)
// 删除文章
export const deleteComment = (id: number) => request('/api/comment/' + id, 'DELETE', null, null)

// 查询分类
export const queryCategory = (
    name: string | null,
    pageNum: number,
    pageSize: number,
) => request('/api/category', 'GET', {
    name: name,
    pageNum: pageNum,
    pageSize: pageSize,
}, null)
// 查询分类选项
export const queryCategoryOptions = () => queryCategory(null, 1, 0x7fffffff).then((data) => {
    return data.data.list.map((item: any) => {
        return {
            value: item.id,
            label: item.name,
        }
    })
})
// 删除分类
export const deleteCategory = (id: number) => request('/api/category/' + id, 'DELETE', null, null)

// 查询标签
export const queryTag = (
    name: string | null,
    pageNum: number,
    pageSize: number,
) => request('/api/tag', 'GET', {
    name: name,
    pageNum: pageNum,
    pageSize: pageSize,
}, null)
// 删除标签
export const deleteTag = (id: number) => request('/api/tag/' + id, 'DELETE', null, null)

// 查询文件
export const queryUpload = (
    displayName: string | null,
    type: string | null,
    status: number | null,
    pageNum: number,
    pageSize: number,
) => request('/api/upload', 'GET', {
    name: displayName,
    type: type,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
}, null)
// 删除文件
export const deleteUpload = (id: number) => request('/api/upload/' + id, 'DELETE', null, null)

// 查询友链
export const queryLink = (
    title: string | null,
    url: string | null,
    status: number | null,
    pageNum: number,
    pageSize: number
) => request('/api/link', 'GET', {
    title: title,
    url: url,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
}, null)
// 删除友链
export const deleteLink = (id: number) => request('/api/link/' + id, 'DELETE', null, null)

// 查询短链
export const queryUrl = (
    key: string | null,
    url: string | null,
    description: string | null,
    status: number | null,
    pageNum: number,
    pageSize: number
) => request('/api/url', 'GET', {
    key: key,
    url: url,
    description: description,
    status: status,
    pageNum: pageNum,
    pageSize: pageSize,
}, null)
// 删除短链
export const deleteUrl = (id: number) => request('/api/url/' + id, 'DELETE', null, null)

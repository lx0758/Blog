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
export const queryComment = (pageNum: number, pageSize: number) => request('/api/comment', 'GET', {
    pageNum: pageNum,
    pageSize: pageSize,
}, null)

// 查询分类
export const queryCategory = () => request('/api/category', 'GET', null, null)

// 查询标签
export const queryTag = () => request('/api/tag', 'GET', null, null)

// 查询文件
export const queryUpload = (pageNum: number, pageSize: number) => request('/api/upload', 'GET', {
    pageNum: pageNum,
    pageSize: pageSize,
}, null)

// 查询短链
export const queryUrl = (pageNum: number, pageSize: number) => request('/api/url', 'GET', {
    pageNum: pageNum,
    pageSize: pageSize,
}, null)

function Comment(options) {
    let element = options['element'] || null;
    let api = new CommentApi(options['apiUrl'], options['apiSubjectId']);

    let init = function () {
        initPanel();
        bindEvent();

        Comment.page.refresh()
    }

    let initPanel = function () {
        let nodeList = document.querySelectorAll(element);
        element = element instanceof HTMLElement ? element : (nodeList[nodeList.length - 1] || null);
        element.classList.add('v')
        element.setAttribute('data-class', 'v');

        element.innerHTML = `
<div class="vpanel">
    <div class="vwrap">
        <p class="cancel-reply text-right" style="display:none;" title="取消回复">
            <svg class="vicon cancel-reply-btn" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4220" width="22" height="22">
                <path
                    d="M796.454 985H227.545c-50.183 0-97.481-19.662-133.183-55.363-35.7-35.701-55.362-83-55.362-133.183V227.545c0-50.183 19.662-97.481 55.363-133.183 35.701-35.7 83-55.362 133.182-55.362h568.909c50.183 0 97.481 19.662 133.183 55.363 35.701 35.702 55.363 83 55.363 133.183v568.909c0 50.183-19.662 97.481-55.363 133.183S846.637 985 796.454 985zM227.545 91C152.254 91 91 152.254 91 227.545v568.909C91 871.746 152.254 933 227.545 933h568.909C871.746 933 933 871.746 933 796.454V227.545C933 152.254 871.746 91 796.454 91H227.545z"
                    p-id="4221">
                </path>
                <path
                    d="M568.569 512l170.267-170.267c15.556-15.556 15.556-41.012 0-56.569s-41.012-15.556-56.569 0L512 455.431 341.733 285.165c-15.556-15.556-41.012-15.556-56.569 0s-15.556 41.012 0 56.569L455.431 512 285.165 682.267c-15.556 15.556-15.556 41.012 0 56.569 15.556 15.556 41.012 15.556 56.569 0L512 568.569l170.267 170.267c15.556 15.556 41.012 15.556 56.569 0 15.556-15.556 15.556-41.012 0-56.569L568.569 512z"
                    p-id="4222">
                </path>
            </svg>
        </p>
        <div class="vheader item3">
            <input name="nick" placeholder="昵称" class="vnick vinput" type="text"><input name="mail" placeholder="邮箱" class="vmail vinput" type="email"><input name="link" placeholder="网址(http://)" class="vlink vinput" type="text">
        </div>
        <div class="vedit">
            <textarea id="veditor" class="veditor vinput" placeholder="赶紧来吐槽一下吧~!"></textarea>
            <div class="vrow">
                <div class="vcol vcol-60 status-bar"></div>
                <div class="vcol vcol-40 vctrl text-right">
                    <span title="表情" class="vicon vemoji-btn">
                        <svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="16172" width="22" height="22">
                            <path
                                d="M512 1024a512 512 0 1 1 512-512 512 512 0 0 1-512 512zM512 56.888889a455.111111 455.111111 0 1 0 455.111111 455.111111 455.111111 455.111111 0 0 0-455.111111-455.111111zM312.888889 512A85.333333 85.333333 0 1 1 398.222222 426.666667 85.333333 85.333333 0 0 1 312.888889 512z"
                                p-id="16173">
                            </path>
                            <path
                                d="M512 768A142.222222 142.222222 0 0 1 369.777778 625.777778a28.444444 28.444444 0 0 1 56.888889 0 85.333333 85.333333 0 0 0 170.666666 0 28.444444 28.444444 0 0 1 56.888889 0A142.222222 142.222222 0 0 1 512 768z"
                                p-id="16174">
                            </path>
                            <path
                                d="M782.222222 391.964444l-113.777778 59.733334a29.013333 29.013333 0 0 1-38.684444-10.808889 28.444444 28.444444 0 0 1 10.24-38.684445l113.777778-56.888888a28.444444 28.444444 0 0 1 38.684444 10.24 28.444444 28.444444 0 0 1-10.24 36.408888z"
                                p-id="16175">
                            </path>
                            <path
                                d="M640.568889 451.697778l113.777778 56.888889a27.875556 27.875556 0 0 0 38.684444-10.24 27.875556 27.875556 0 0 0-10.24-38.684445l-113.777778-56.888889a28.444444 28.444444 0 0 0-38.684444 10.808889 28.444444 28.444444 0 0 0 10.24 38.115556z"
                                p-id="16176">
                            </path>
                        </svg>
                    </span>
                    <span title="预览" class="vicon vpreview-btn">
                        <svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="17688" width="22" height="22">
                            <path
                                d="M502.390154 935.384615a29.538462 29.538462 0 1 1 0 59.076923H141.430154C79.911385 994.461538 29.538462 946.254769 29.538462 886.153846V137.846154C29.538462 77.745231 79.950769 29.538462 141.390769 29.538462h741.218462c61.44 0 111.852308 48.206769 111.852307 108.307692v300.268308a29.538462 29.538462 0 1 1-59.076923 0V137.846154c0-26.899692-23.355077-49.230769-52.775384-49.230769H141.390769c-29.420308 0-52.775385 22.331077-52.775384 49.230769v748.307692c0 26.899692 23.355077 49.230769 52.775384 49.230769h360.999385z"
                                p-id="17689">
                            </path>
                            <path
                                d="M196.923077 216.615385m29.538461 0l374.153847 0q29.538462 0 29.538461 29.538461l0 0q0 29.538462-29.538461 29.538462l-374.153847 0q-29.538462 0-29.538461-29.538462l0 0q0-29.538462 29.538461-29.538461Z"
                                p-id="17690">
                            </path>
                            <path
                                d="M649.846154 846.769231a216.615385 216.615385 0 1 0 0-433.230769 216.615385 216.615385 0 0 0 0 433.230769z m0 59.076923a275.692308 275.692308 0 1 1 0-551.384616 275.692308 275.692308 0 0 1 0 551.384616z"
                                p-id="17691">
                            </path>
                            <path
                                d="M807.398383 829.479768m20.886847-20.886846l0 0q20.886846-20.886846 41.773692 0l125.321079 125.321079q20.886846 20.886846 0 41.773693l0 0q-20.886846 20.886846-41.773693 0l-125.321078-125.321079q-20.886846-20.886846 0-41.773693Z"
                                p-id="17692">
                            </path>
                        </svg>
                    </span>
                </div>
            </div>
        </div>
        <div class="vrow">
            <div class="vcol vcol-30">
                <a alt="Markdown is supported" href="https://guides.github.com/features/mastering-markdown/"
                    class="vicon" target="_blank">
                    <svg class="markdown" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true">
                        <path
                            d="M14.85 3H1.15C.52 3 0 3.52 0 4.15v7.69C0 12.48.52 13 1.15 13h13.69c.64 0 1.15-.52 1.15-1.15v-7.7C16 3.52 15.48 3 14.85 3zM9 11H7V8L5.5 9.92 4 8v3H2V5h2l1.5 2L7 5h2v6zm2.99.5L9.5 8H11V5h2v3h1.5l-2.51 3.5z"
                            fill-rule="evenodd">
                        </path>
                    </svg>
                </a>
            </div>
            <div class="vcol vcol-70 text-right">
                <button type="button" title="Cmd|Ctrl+Enter" class="vsubmit vbtn">提交</button>
            </div>
        </div>
        <div class="vemojis" style="display:none;">
            <i>😀</i>
            <i>😃</i>
            <i>😄</i>
            <i>😁</i>
            <i>😆</i>
            <i>😅</i>
            <i>😂</i>
            <i>😊</i>
            <i>😇</i>
            <i>😉</i>
            <i>😌</i>
            <i>😍</i>
            <i>😘</i>
            <i>😗</i>
            <i>😙</i>
            <i>😚</i>
            <i>😋</i>
            <i>😜</i>
            <i>😝</i>
            <i>😛</i>
            <i>😎</i>
            <i>😏</i>
            <i>😒</i>
            <i>😞</i>
            <i>😔</i>
            <i>😟</i>
            <i>😕</i>
            <i>😣</i>
            <i>😖</i>
            <i>😫</i>
            <i>😩</i>
            <i>😠</i>
            <i>😡</i>
            <i>😶</i>
            <i>😐</i>
            <i>😑</i>
            <i>😯</i>
            <i>😦</i>
            <i>😧</i>
            <i>😮</i>
            <i>😲</i>
            <i>😵</i>
            <i>😳</i>
            <i>😱</i>
            <i>😨</i>
            <i>😰</i>
            <i>😢</i>
            <i>😥</i>
            <i>😭</i>
            <i>😓</i>
            <i>😪</i>
            <i>😴</i>
            <i>😷</i>
            <i>😈</i>
            <i>😺</i>
            <i>😸</i>
            <i>😹</i>
            <i>😻</i>
            <i>😼</i>
            <i>😽</i>
            <i>🙀</i>
            <i>😿</i>
            <i>😾</i>
            <i>🐱</i>
            <i>🐭</i>
            <i>🐮</i>
            <i>🐵</i>
            <i>✋</i>
            <i>✊</i>
            <i>✌️</i>
            <i>👆</i>
            <i>👇</i>
            <i>👈</i>
            <i>👉</i>
            <i>👊</i>
            <i>👋</i>
            <i>👏</i>
            <i>👐</i>
            <i>👍</i>
            <i>👎</i>
            <i>👌</i>
            <i>🙏</i>
            <i>👂</i>
            <i>👀</i>
            <i>👃</i>
            <i>👄</i>
            <i>👅</i>
            <i>❤️</i>
            <i>💘</i>
            <i>💖</i>
            <i>⭐️</i>
            <i>✨</i>
            <i>⚡️</i>
            <i>☀️</i>
            <i>☁️</i>
            <i>❄️</i>
            <i>☔️</i>
            <i>☕️</i>
            <i>✈️</i>
            <i>⚓️</i>
            <i>⌚️</i>
            <i>☎️</i>
            <i>⌛️</i>
            <i>✉️</i>
            <i>✂️</i>
            <i>✒️</i>
            <i>✏️</i>
            <i>❌</i>
            <i>♻️</i>
            <i>✅</i>
            <i>❎</i>
            <i>Ⓜ️</i>
            <i>ℹ️</i>
            <i>™️</i>
            <i>©️</i>
            <i>®️</i>
        </div>
        <div style="display:none;" class="vmark"></div>
    </div>
</div>
<div class="vcount" style="display: block;">
    共
    <span class="vnum">0</span>
    条评论
</div>
<div class="vload-top text-center" style="display:none;">
    <i class="vspinner" style="width:30px;height:30px;"></i>
</div>
<div class="vcards"></div>
<div class="vload-bottom text-center" style="display: none;">
    <i class="vspinner" style="width:30px;height:30px;"></i>
</div>
<div class="vempty" style="display:none;">赶紧来抢个沙发~!</div>
<div class="vpage txt-center" style="display: block;">
    <button type="button" class="vmore vbtn">加载更多...</button>
</div>`;
    }

    let bindEvent = function () {
        let _vemojis = utils.find(element, '.vemojis');
        let _emojiCtrl = utils.find(element, '.vemoji-btn');
        let _emojiIcon = utils.find(element, '.vemoji-btn>svg');
        Comment.emoji = {
            show() {
                utils.attr(_emojiCtrl, 'v', 1);
                utils.attr(_vemojis, 'style', 'display:block');
                utils.attr(_emojiIcon, 'fill', '#ef2f11');
            },
            hide() {
                utils.removeAttr(_emojiCtrl, 'v');
                utils.attr(_vemojis, 'style', 'display:hide');
                utils.attr(_emojiIcon, 'fill', '');
            },
        }
        let _emojiClick = function (element, value) {
            if (document.selection) {
                //For browsers like Internet Explorer
                element.focus();
                let sel = document.selection.createRange();
                sel.text = value;
                element.focus();
            } else if (element.selectionStart || element.selectionStart == '0') {
                //For browsers like Firefox and Webkit based
                let startPos = element.selectionStart;
                let endPos = element.selectionEnd;
                let scrollTop = element.scrollTop;
                element.value = element.value.substring(0, startPos) + value + element.value.substring(endPos, element.value.length);
                element.focus();
                element.selectionStart = startPos + value.length;
                element.selectionEnd = startPos + value.length;
                element.scrollTop = scrollTop;
            } else {
                element.focus();
                element.value += value;
            }
        }
        let nodes = _vemojis.childNodes
        for (let i = 0; i < nodes.length; i++) {
            let item = nodes.item(i)
            utils.on('click', item, (e) => {
                _emojiClick(_content, item.innerHTML)
            });
        }
        utils.on('click', _emojiCtrl, (e) => {
            let _vi = utils.attr(_emojiCtrl, 'v');
            if (_vi) Comment.emoji.hide()
            else Comment.emoji.show();
        });


        let _panel = utils.find(element, '.vpanel');
        let _wrap = utils.find(element, '.vwrap');
        let _cancel = utils.find(element, '.cancel-reply');
        let _nickname = utils.find(element, '.vnick');
        let _email = utils.find(element, '.vmail');
        let _link = utils.find(element, '.vlink');
        let _content = utils.find(element, '.veditor');
        let _status = utils.find(element, '.status-bar');
        let _submit = utils.find(element, '.vsubmit');
        Comment.panel = {
            replyAt(newCard, comment) {
                let _wrapper = utils.find(newCard, '.vreply-wrapper')
                if (_wrapper) {
                    utils.attr(_wrap, 'parent-id', comment.id);
                    utils.attr(_cancel, 'style', 'display:block');
                    _wrap.parentNode.removeChild(_wrap)
                    _wrapper.appendChild(_wrap)
                    utils.attr(_content, 'placeholder', '@' + comment.nickname)
                    _content.focus()
                }
            },
            replyCancel() {
                utils.attr(_wrap, 'parent-id', null);
                utils.attr(_cancel, 'style', 'display:none');
                _wrap.parentNode.removeChild(_wrap)
                _panel.appendChild(_wrap)
                utils.attr(_content, 'placeholder', '赶紧来吐槽一下吧~!')
                _content.focus()
            },
            showInfo(info) {
                _status.innerHTML = info
                setTimeout(function () {
                    _status.innerHTML = null
                },3000);
            }
        }
        utils.on('click', _cancel, (e) => {
            Comment.panel.replyCancel()
        });
        utils.on('click', _submit, (e) => {
            let parentId = utils.attr(_wrap, 'parent-id')
            let nickname = _nickname.value
            let email = _email.value
            let link = _link.value
            let content = _content.value

            if (!nickname || nickname.length === 0) {
                Comment.panel.showInfo('昵称不能为空')
                return
            }
            if (!content || content.length === 0) {
                Comment.panel.showInfo('评论内容不能为空')
                return
            }

            api.commit(parentId, nickname, email, link, content).then(data => {
                _nickname.value = null
                _email.value = null
                _link.value = null
                _content.value = null
                Comment.panel.showInfo('评论成功!')
                Comment.panel.replyCancel()
                Comment.emoji.hide()
                Comment.page.refresh()
            }).catch(function (error) {
                Comment.panel.showInfo(error)
            })
        });

        let _total = utils.find(element, '.vnum');
        let _count = utils.find(element, '.vcount');
        let _empty = utils.find(element, '.vempty');
        let _loadingTop = utils.find(element, '.vload-top');
        let _loadingBottom = utils.find(element, '.vload-bottom');
        let _loadMore = utils.find(element, '.vpage');
        let _loadMoreBtn = utils.find(element, '.vmore');
        Comment.page = {
            refresh() {
                utils.attr(_count, 'style', 'display:none');
                utils.attr(_empty, 'style', 'display:none');
                utils.attr(_loadingTop, 'style', 'display:block');
                utils.attr(_loadMore, 'style', 'display:none');
                api.refresh().then(data => {
                    let total = data.total
                    let hasMore = data.hasMore
                    let comments = data.comments

                    if (comments && comments.length > 0) {
                        utils.find(element, '.vcards').innerHTML = ''
                        comments.forEach(comment => {
                            insertComment(comment)
                        })
                        Comment.page.refreshFinish(total, hasMore)
                    } else {
                        Comment.page.refreshFinish(0, false)
                    }
                }).catch(function (err) {
                    Comment.page.refreshFinish(0, false)
                })
            },
            refreshFinish(total, hasMore) {
                _total.innerHTML = total
                if (total > 0) {
                    utils.attr(_count, 'style', 'display:block');
                    utils.attr(_empty, 'style', 'display:none');
                } else {
                    utils.attr(_count, 'style', 'display:none');
                    utils.attr(_empty, 'style', 'display:block');
                }
                utils.attr(_loadingTop, 'style', 'display:none');
                this.loadMoreFinish(hasMore)
            },
            loadMore() {
                utils.attr(_loadingBottom, 'style', 'display:block');
                utils.attr(_loadMore, 'style', 'display:none');
                api.loadMore().then(data => {
                    let total = data.total
                    let hasMore = data.hasMore
                    let comments = data.comments

                    Comment.page.loadMoreFinish(hasMore)
                    if (comments && comments.length > 0) {
                        comments.forEach(comment => {
                            insertComment(comment)
                        })
                    }
                }).catch(function (error) {
                    Comment.page.loadMoreFinish(false)
                })
            },
            loadMoreFinish(hasMore) {
                utils.attr(_loadingBottom, 'style', 'display:none');
                if (hasMore) {
                    utils.attr(_loadMore, 'style', 'display:block');
                } else {
                    utils.attr(_loadMore, 'style', 'display:none');
                }
            },
        }
        utils.on('click', _loadMoreBtn, (e) => {
            Comment.page.loadMore()
        });
    }

    let insertComment = function (comment) {
        let newCard = utils.create('div', {'class': 'vcard'})

        newCard.innerHTML = `
<img class="vimg" src="${comment.avatar}">
<div class="vh">
    <div class="vhead">
        <span class="vnick">${comment.nickname}</span>
        <span class="vsys">${comment.browser}</span>
        <span class="vsys">${comment.system}</span>
    </div>
    <div class="vmeta">
        <span class="vtime">${comment.time}</span>
        <span class="vat" data-vm-id="${comment.id}" data-self-id="${comment.id}">回复</span>
    </div>
    <pre class="vcontent" data-expand="查看更多...">${comment.content}</pre>
    <div class="vquote" data-self-id="${comment.id}"></div>
    <div class="vreply-wrapper" data-self-id="${comment.id}"></div>
</div>`

        comment.children.forEach(childComment => {
            let newChildCard = utils.create('div', {'class': 'vcard', 'id': childComment.id})
            newChildCard.innerHTML = `
<img class="vimg" src="${comment.avatar}">
<div class="vh">
    <div class="vhead">
        <span class="vnick">${childComment.nickname}</span>
        <span class="vsys">${childComment.browser}</span>
        <span class="vsys">${childComment.system}</span>
    </div>
    <div class="vmeta">
        <span class="vtime">${childComment.time}</span>
    </div>
    <pre class="vcontent" data-expand="查看更多...">${childComment.content}</pre>
</div>`
            utils.find(newCard, '.vquote').appendChild(newChildCard)
        })

        let _replay = utils.find(newCard, '.vat')
        utils.on('click', _replay, (e) => {
            Comment.panel.replyAt(newCard, comment)
        })

        let _list = utils.find(element, '.vcards')
        _list.appendChild(newCard)
    }

    init()
}

function CommentApi(apiUrl, apiSubjectId) {
    let url = apiUrl;
    let subjectId = apiSubjectId;

    let refresh = function () {
        page = 1
        return query(page)
    }

    let loadMore = function () {
        page += 1
        return query(page)
    }

    let page
    let query = function (page) {
        return new Promise(function (resolve, reject) {
            // let example = {
            //     total: 0,
            //     hasMore: true,
            //     comments: [
            //         {
            //             id: 0,
            //             avatar: "https://gravatar.loli.net/avatar/d41d8cd98f00b204e9800998ecf8427e?d=mp&amp;v=1.4.14",
            //             nickname: "张三",
            //             browser: "Chrome 88.0.4324.182",
            //             system: "Windows 10.0",
            //             time: '2021-3-2 11:32:51',
            //             content: '我是评论内容',
            //             children: []
            //         },
            //     ]
            // }
            $.ajax({
                url: url,
                type: 'GET',
                data: {
                    subjectId: subjectId,
                    page: page,
                },
                success: function (response) {
                    if (response.status === 200) {
                        resolve(response.data)
                    } else {
                        reject(response.message)
                    }
                },
                error: function (XMLHttpRequest, textStatus, errorThrown) {
                    reject('网络请求失败')
                },
                complete: function(XMLHttpRequest, textStatus) {

                }
            });
        });
    }

    let commit = function (parentId, nickname, email, link, content) {
        return new Promise(function (resolve, reject) {
            $.ajax({
                url: url,
                type: 'POST',
                data: {
                    subjectId: subjectId,
                    parentId: parentId,
                    nickname: nickname,
                    email: email,
                    link: link,
                    content: content,
                },
                success: function (response) {
                    if (response.status === 200) {
                        resolve()
                    } else {
                        reject(response.message)
                    }
                },
                error: function (XMLHttpRequest, textStatus, errorThrown) {
                    reject('网络请求失败')
                },
                complete: function (XMLHttpRequest, textStatus) {

                }
            });
        });
    }

    return {
        refresh: refresh,
        loadMore: loadMore,
        commit: commit,
    }
}

const unescapeMap = {};
const escapeMap = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#39;',
    '`': '&#x60;',
    "\\": '&#x5c;'
};
for (let key in escapeMap) {
    unescapeMap[escapeMap[key]] = key;
}

const reUnescapedHtml = /[&<>"'`\\]/g
const reHasUnescapedHtml = RegExp(reUnescapedHtml.source)
const reEscapedHtml = /&(?:amp|lt|gt|quot|#39|#x60|#x5c);/g
const reHasEscapedHtml = RegExp(reEscapedHtml.source)

const utils = {
    on(type, el, handler, capture) {
        type = type.split(' ')
        for (let i = 0, len = type.length; i < len; i++) {
            utils.off(type[i], el, handler, capture)
            if (el.addEventListener) el.addEventListener(type[i], handler, capture || false);
            else if (el.attachEvent) el.attachEvent(`on${type[i]}`, handler);
            else el[`on${type[i]}`] = handler;
        }
    },
    off(type, el, handler, capture) {
        type = type.split(' ')
        for (let i = 0, len = type.length; i < len; i++) {
            if (el.removeEventListener) el.removeEventListener(type, handler, capture || false);
            else if (el.detachEvent) el.detachEvent(`on${type}`, handler);
            else el[`on${type}`] = null;
        }
    },

    escape(s) {
        return (s && reHasUnescapedHtml.test(s)) ?
            s.replace(reUnescapedHtml, (chr) => escapeMap[chr]) :
            s
    },

    unescape(s) {
        return (s && reHasEscapedHtml.test(s)) ?
            s.replace(reEscapedHtml, (entity) => unescapeMap[entity]) :
            s
    },
    /**
     * Create Element
     * @param {String} name ElementTagName
     * @param {Object} attrName
     * @param {Object} attrVal
     */
    create(name, attrName, attrVal) {
        let el = document.createElement(name)
        utils.attr(el, attrName, attrVal)
        return el
    },
    /**
     * el.querySelector
     * @param {HTMLElement} el HTMLElement
     * @param {String} selector
     */
    find(el, selector) {
        return el.querySelector(selector)
    },

    /**
     * el.querySelectorAll
     * @param {HTMLElement} el HTMLElement
     * @param {String} selector
     */
    findAll(el, selector) {
        return el.querySelectorAll(selector)
    },

    /**
     * get/set attributes
     * @param {HTMLElement} el
     * @param {String | Object} name
     * @param {String} value
     */
    attr(el, name, value) {
        if (typeof el.getAttribute === "undefined") return utils.prop(el, name, value)
        if (value !== undefined) {
            if (value === null) utils.removeAttr(el, name)
            else el.setAttribute(name, value)
        } else if (({}).toString.call(name) === '[object Object]') {
            utils.each(name, (k, v) => {
                el.setAttribute(k, v)
            })
        } else return el.getAttribute(name)
    },
    /**
     * get/set props
     * @param {HTMLElement} el
     * @param {String} name
     * @param {String} value
     */
    prop(el, name, value) {
        if (value !== undefined) return el[name] = value
        else if (({}).toString.call(name) === '[object Object]') {
            utils.each(name, (k, v) => {
                el[k] = v
            })
        } else return el[name]
    },
    /**
     * Remove el attribute
     * @param {HTMLElement} el
     * @param {String} names
     * @returns {HTMLElement} el
     */
    removeAttr(el, names) {
        let name,
            i = 0,
            // Attribute names can contain non-HTML whitespace characters
            // https://html.spec.whatwg.org/multipage/syntax.html#attributes-2
            attrNames = names && names.match(/[^\x20\t\r\n\f\*\/\\]+/g);
        if (attrNames && el.nodeType === 1) {
            while ((name = attrNames[i++])) {
                el.removeAttribute(name);
            }
        }
        return el
    },
    /**
     * Clear element attributes
     * @param {HTMLElement} el
     */
    clearAttr(el) {
        let attrs = el.attributes
        let ignoreAttrs = ['align', 'alt', 'checked', 'class', 'disabled', 'href', 'id', 'target', 'title', 'type', 'src', 'style']
        utils.each(attrs, (idx, attr) => {
            let name = attr.name
            switch (name.toLowerCase()) {
                case 'style':
                    let style = attr.value
                    utils.each(style.split(';'), (idx, item) => {
                        if (item.indexOf('color') > -1) utils.attr(el, 'style', item);
                        else utils.removeAttr(el, 'style');
                    })
                    break;
                case 'class':
                    if (el.nodeName == 'CODE') return false
                    let clazz = attr.value
                    if (clazz.indexOf('at') > -1) utils.attr(el, 'class', 'at');
                    else if (clazz.indexOf('vemoji') > -1) utils.attr(el,'class','vemoji');
                    else utils.removeAttr(el, 'class')
                    break;
                default:
                    if (ignoreAttrs.indexOf(name) > -1) return true
                    else utils.removeAttr(el, name)
                    break;

            }

        })
        return el
    },
    /**
     * Remove Child node
     * @param {HTMLElement} child
     */
    remove(child) {
        try {
            if (child.parentNode) child.parentNode.removeChild(child)
        } catch (error) {}
    },

    /**
     * 遍历对象或者数组
     * collection, callback(indexInArray, valueOfElement)
     * @param {Object} collection
     * @param {Function} callback
     * @return {Object} collection
     */
    each(collection, callback) {
        let value,
            i = 0,
            length = collection.length,
            likeArray = ["[object Array]", "[object NodeList]"],
            type = ({}).toString.call(collection)
        if (likeArray.indexOf(type) > -1) {
            for (; i < length; i++) {
                if (callback && callback.call(collection[i], i, collection[i]) === false) break
            }
        } else {
            for (i in collection) {
                if (collection.hasOwnProperty(i)) {
                    if (callback && callback.call(collection[i], i, collection[i]) === false) break
                }
            }
        }
        return collection
    }
}

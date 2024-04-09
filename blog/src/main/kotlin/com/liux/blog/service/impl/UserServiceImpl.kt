package com.liux.blog.service.impl

import com.liux.blog.bean.UserDetailsImpl
import com.liux.blog.bean.event.BaseInfoUpdateEvent
import com.liux.blog.bean.po.User
import com.liux.blog.dao.UserMapper
import com.liux.blog.service.UserService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.ApplicationEventPublisher
import org.springframework.security.core.userdetails.UserDetails
import org.springframework.security.core.userdetails.UsernameNotFoundException
import org.springframework.security.crypto.password.PasswordEncoder
import org.springframework.stereotype.Service
import java.util.*

@Service
class UserServiceImpl: UserService {

    @Autowired
    private lateinit var applicationEventPublisher: ApplicationEventPublisher

    @Autowired
    private lateinit var passwordEncoder: PasswordEncoder

    @Autowired
    private lateinit var userMapper: UserMapper

    @Throws(UsernameNotFoundException::class)
    override fun loadUserByUsername(username: String): UserDetails {
        val user = userMapper.getByUsername(username) ?: throw UsernameNotFoundException("Account does not exist")
        return UserDetailsImpl(user, emptyList())
    }

    override fun getByOwner(): User {
        return userMapper.getByOwner()
    }

    override fun getById(id: Int): User? {
        return userMapper.getByPrimaryKey(id)
    }

    override fun listByAdmin(username: String?, nickname: String?, status: Int?): List<User> {
        return userMapper.selectByAdmin(User(
            username = username,
            nickname = nickname,
            status = status,
        ))
    }

    override fun updateByLogin(id: Int) {
        userMapper.updateByPrimaryKeySelective(User(
            id = id,
            lastLoginTime = Date(),
        ))
    }

    override fun updateBySelf(
        id: Int,
        avatar: String?,
        nickname: String,
        description: String?,
        email: String,

        github: String?,
        weibo: String?,
        google: String?,
        twitter: String?,
        facebook: String?,
        stackOverflow: String?,
        youtube: String?,
        instagram: String?,
        skype: String?,
    ): Int {
        val accounts = User.Accounts(
            github = github,
            weibo = weibo,
            google = google,
            twitter = twitter,
            facebook = facebook,
            stackOverflow = stackOverflow,
            youtube = youtube,
            instagram = instagram,
            skype = skype,
        )
        val user = User(
            id = id,
            avatar = avatar,
            nickname = nickname,
            description = description,
            email = email,
            accounts = accounts,
            updateTime = Date(),
        )
        val updateRow = userMapper.updateByPrimaryKeyNullable(user)
        if (updateRow > 0) {
            applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.USER)
        }
        return updateRow
    }

    override fun verifyPassword(id: Int, password: String): Boolean {
        if (password.isEmpty()) return false
        val user = getById(id)!!
        return passwordEncoder.matches(password, user.password)
    }

    override fun updatePassword(id: Int, password: String): Boolean {
        val encodedPassword = passwordEncoder.encode(password)
        val updateRow = userMapper.updateByPrimaryKeySelective(User(
            id = id,
            password = encodedPassword,
            updateTime = Date(),
        ))
        return updateRow > 0
    }
}

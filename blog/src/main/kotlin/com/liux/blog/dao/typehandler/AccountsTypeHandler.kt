package com.liux.blog.dao.typehandler

import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.liux.blog.bean.po.User
import org.apache.ibatis.type.BaseTypeHandler
import org.apache.ibatis.type.JdbcType
import org.apache.ibatis.type.MappedTypes
import java.sql.CallableStatement
import java.sql.PreparedStatement
import java.sql.ResultSet

@MappedTypes(User.Accounts::class)
class AccountsTypeHandler : BaseTypeHandler<User.Accounts>() {

    override fun setNonNullParameter(ps: PreparedStatement?, index: Int, parameter: User.Accounts?, jdbcType: JdbcType?) {
        ps?.setString(index, jacksonObjectMapper().writeValueAsString(parameter) ?: "{}")
    }

    override fun getNullableResult(rs: ResultSet?, columnName: String): User.Accounts? {
        val json = rs?.getString(columnName)
        return json2Object(json)
    }

    override fun getNullableResult(rs: ResultSet?, columnIndex: Int): User.Accounts? {
        val json = rs?.getString(columnIndex)
        return json2Object(json)
    }

    override fun getNullableResult(cs: CallableStatement?, columnIndex: Int): User.Accounts? {
        val json = cs?.getString(columnIndex)
        return json2Object(json)
    }

    private fun json2Object(json: String?): User.Accounts? {
        if (json == null) return null
        return jacksonObjectMapper().readValue(json, User.Accounts::class.java)
    }
}
package com.liux.blog.dao

import com.liux.blog.bean.po.Fragment
import org.apache.ibatis.annotations.Param
import org.springframework.stereotype.Repository

@Repository
interface FragmentMapper {
    fun deleteByPrimaryKey(id: Int?): Int

    fun insert(record: Fragment?): Int

    fun insertSelective(record: Fragment?): Int

    fun getByPrimaryKey(id: Int?): Fragment?

    fun getByIdOrKey(@Param("id") id: Int, @Param("key") key: String): Fragment?

    fun selectByAdmin(fragment: Fragment): List<Fragment>

    fun updateByPrimaryKeySelective(record: Fragment?): Int

    fun updateByPrimaryKey(record: Fragment?): Int
}
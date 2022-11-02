<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 13:50
 **/
namespace Design\Adapter\Database;


interface Datebase {

    /**
     * 连接
     */
    function connect ($host,$port);

    /**
     * 查询
     */
    function query();

}
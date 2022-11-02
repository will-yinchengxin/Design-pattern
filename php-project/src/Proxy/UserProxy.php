<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 17:03
 **/
namespace Design\Proxy;

class UserProxy implements Proxy {

    /**
     * 从库只读
     */
    public function getName(int $id)
    {
        $db = DB::getDateBase('slave');
        return $db->query("select * from user where id = {$id}");
    }

    /**
     * 主库写
     */
    public function updateName(int $id, string $name)
    {
        $db = DB::getDateBase('master');
        return $db->query("update user set name = {$name} where user = {$id}");
    }

}
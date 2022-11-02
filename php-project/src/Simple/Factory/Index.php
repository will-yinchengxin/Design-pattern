<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 11:56
 **/
namespace Design\Simple\Factory;

use Design\Simple\Register\Index as Register;


class Index {

    public static function getData()
    {
        return (new date())->index();
    }

    public function setRegister()
    {
        Register::set('data', self::getData());
    }
}
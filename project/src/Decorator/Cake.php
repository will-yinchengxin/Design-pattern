<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 16:10
 **/
namespace Design\Decorator;

class Cake implements Food {

    public function make()
    {
        echo "制作蛋糕";
    }

}
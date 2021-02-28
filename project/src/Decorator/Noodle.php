<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 16:11
 **/
namespace Design\Decorator;

class Noodle implements Food {

    public function make()
    {
        echo "制作面条";
    }

}
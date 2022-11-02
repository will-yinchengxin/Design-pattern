<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 17:02
 **/
namespace Design\Proxy;

interface Proxy {

    public function getName(int $id);

    public function updateName(int $id, string $name);

}
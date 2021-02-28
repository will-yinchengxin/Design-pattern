<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 16:52
 **/
namespace Design\Iterator;

class Index implements \Iterator {

    protected $index = 0;

    protected $data = [];

    public function __construct()
    {
     $link = mysqli_connect('localhost', 'root', '123', 'market');
     $rec = mysqli_query($link, 'select id from user');
     $this->data = mysqli_fetch_all($rec, MYSQLI_ASSOC);
    }

    /**
     * 获取当前内容
     */
    public function current()
    {
        $id = $this->data[$this->index];
        return User::find($id);
    }

    /**
     * 移动key到下一个
     */
    public function next()
    {
        return $this->index++;
    }

    /**
     * 获取迭代器key位置
     */
    public function key()
    {
        return $this->index;
    }

    /**
     * 校验迭代器是否有数据
     */
    public function valid()
    {
        return $this->index < count($this->data);
    }

    /**
     * 重置迭代器
     */
    public function rewind()
    {
        $this->index = 0;
    }

}
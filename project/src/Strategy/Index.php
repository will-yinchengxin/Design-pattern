<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 14:11
 **/
namespace Design\Strategy;

class Index {

    /**
     * @var $strategy
     */
    protected $strategy ;

    public function page()
    {
        echo "this is index page";
        // 展示广告
        $this->strategy->showAd();
        // 展示分类
        $this->strategy->showCategory();
    }

    /**
     * 设置传输类型,减少判断
     */
    public function setStrategy(Strategy $strategy)
    {
        $this->strategy = $strategy;
    }

    public function judgeStrategy() {
        if ($_GET['man']) {
            $strategy = new ManStrategy();
        } elseif ($_GET['woman']) {
            $strategy = new WomanStrategy();
        }
        $this->setStrategy($strategy);

        return $this;
    }

}

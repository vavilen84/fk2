<?php

namespace App\Tests\Functional;

use \FunctionalTester;
use App\Tests\Functional\BaseFunctionalCest;

class HelloFromTestCest extends BaseFunctionalCest
{
    public function run(FunctionalTester $I)
    {
        $I->wantTo('Run');
    }
}

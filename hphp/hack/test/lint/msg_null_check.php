<?hh // strict
/**
 * Copyright (c) 2014, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *
 */

function takes_int(int $x): void {}

class X {
  private ?int $x = null;

  public function test(): void {
    if($this->x) {
      takes_int($this->x);
    }
  }
}

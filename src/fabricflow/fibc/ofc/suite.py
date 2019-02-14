#! /usr/bin/env python
# -*- coding: utf-8 -*-

# Copyright (C) 2018 Nippon Telegraph and Telephone Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
# implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""
Test Sute
"""

import unittest

from fabricflow.fibc.ofc import ofc_test
from fabricflow.fibc.ofc import generic_test

_PKGS = [
    ofc_test,
    generic_test,
]

def _suite():
    suite = unittest.TestSuite()
    for pkg in _PKGS:
        for tst in pkg.TESTS:
            suite.addTest(unittest.makeSuite(tst))

    return suite

if __name__ == "__main__":
    unittest.TextTestRunner().run(_suite())

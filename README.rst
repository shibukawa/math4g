Math for Game
=================

Math for Game is a Golang library that provides math functions for game.
This code is comes from `nanovgo <https://github.com/shibukawa/nanovgo>`_ and added gopher.js optimization.

This library provides math functions that uses float32.
But, Gopher.js recommends using int or float64 instead of other types
(like float32, uint32 and so on) for performance.
It uses float64 for gopher.js

Document
---------

See `godoc <https://godoc.org/github.com/shibukawa/math4g>`_.

Benchmark
----------

.. list-table::
   :header-rows: 1

   - * Function
     * ``math`` with go
     * ``math4go`` with go
     * ``math`` with gopher.js
     * ``math4go`` with gopher.js
   - * ``Abs``
     * 7.00 ns/op
     * 1.46 ns/op
     * 17.6 ns/op
     * 16.6 ns/op
   - * ``Floor``
     * 5.81 ns/op
     * 1.48 ns/op
     * 19.8 ns/op
     * 19.0 ns/op
   - * ``Ceil``
     * 5.78 ns/op
     * 1.47 ns/op
     * 36.6 ns/op
     * 20.2 ns/op
   - * ``Trunc``
     * 7.07 ns/op
     * 1.47 ns/op
     * 17.0 ns/op
     * 11.2 ns/op
   - * ``IsNaN``
     * 1.49 ns/op
     * 1.48 ns/op
     * 11.0 ns/op
     * 11.0 ns/op
   - * ``NaN``
     * 0.42 ns/op
     * 0.32 ns/op
     * 0.49 ns/op
     * 0.49 ns/op
   - * ``Cbrt``
     * 34.0 ns/op
     * 34.0 ns/op
     * 1374 ns/op
     * 114 ns/op


License
--------

MIT

Author
--------

Yoshiki Shibukawa

Thanks
----------

github.com/rkusa/gm/math32
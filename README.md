# errby

Package `github.com/elblox/errby` is a small helper package to compare errors
by their output with [`is`](https://github.com/matryer/is), the famous lightweight
testing mini-framework by Mat Ryer.

As Dave Cheney points out in [Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully),
the output of `error.Error` should not be inspected in order to change the
behaviour of a program, but

Full quote:

> As an aside, I believe you should never inspect the output of the error.Error
> method. The Error method on the error interface exists for humans, not code.
>
> The contents of that string belong in a log file, or displayed on screen. You
> shouldn’t try to change the behaviour of your program by inspecting it.
>
> I know that sometimes this isn’t possible, and as someone pointed out on
> twitter, this advice doesn’t apply to writing tests. Never the less, comparing
> the string form of an error is, in my opinion, a code smell, and you should
> try to avoid it.

As pointed out in the last section, there are good reasons to do this in tests.

If you find your self in the situation, where you need to compare the output
of errors in your tests, this small helper may become handy.

## Links

* [`is`](https://github.com/matryer/is)
* [`moq`](https://github.com/matryer/moq)

## License

This code is licensed 2021 by Elblox AG under the MIT license
(see [LICENSE](LICENSE) for the full text).

# sha256

After reading the [Federal Information Processing Standards Publication
180-2](https://web.archive.org/web/20130526224224if_/http://csrc.nist.gov/groups/STM/cavp/documents/shs/sha256-384-512.pdf)
I thought I'd have a stab at implementing SHA256 in go, to check I've
understood the algorithm.

This isn't meant to be pretty, I just wanted to have an idea what the algorithm
does (which is still quite a stretch from understanding why it is
cryptographically secure).

The implementation seems to work, so I got it to output its calculations in
the same format as the above paper.

# Building

```
go build -o sha256
```

# Running

```
./sha256 <file>  # to calculate sha256 of a file
./sha256         # to run examples given in FIPS 180-2 paper (pages 10-15)
```

## Future plans

I'd like to implement SHA256 in Z80 assembly for the ZX Spectrum, for no good
reason other than it is a bit of fun. I saw somebody else [did this
already](https://bitcointalk.org/index.php?topic=397014.0), but I want to have
a go at it myself too, and wrap it with a BASIC Loader and create a docker
container that runs it under the Free Unix Spectrum Emulator, so you can
calculate a sha256 on your host computer using an emulated ZX Spectrum. This is
entirely ridiculous, and intentionally so.

I'd like to have a go at writing it also in aarch64 assembly, both with and
without using the ARMv8 Cryptography Extensions (under Advanced SIMD). Maybe as
a userspace program, or a raspberry pi kernel. I'd like to get it to
parallelise the calculation on all available cores, in order to learn how
concurrent programming is performed in aarch64 assembly.

I expect there are many examples on the web, I just wanted to do it myself as a
learning exercise, and have a bit of fun...

## Demo output

This is the output from running `sha256` without arguments, which recreates
[the examples in FIPS 180-2 pages
10-15](https://web.archive.org/web/20130526224224if_/http://csrc.nist.gov/groups/STM/cavp/documents/shs/sha256-384-512.pdf):

### Hash of "abc" (block 1)

```
            a         b         c         d         e         f         g         h

init:   6a09e667  bb67ae85  3c6ef372  a54ff53a  510e527f  9b05688c  1f83d9ab  5be0cd19
t =  0  5d6aebcd  6a09e667  bb67ae85  3c6ef372  fa2a4622  510e527f  9b05688c  1f83d9ab
t =  1  5a6ad9ad  5d6aebcd  6a09e667  bb67ae85  78ce7989  fa2a4622  510e527f  9b05688c
t =  2  c8c347a7  5a6ad9ad  5d6aebcd  6a09e667  f92939eb  78ce7989  fa2a4622  510e527f
t =  3  d550f666  c8c347a7  5a6ad9ad  5d6aebcd  24e00850  f92939eb  78ce7989  fa2a4622
t =  4  04409a6a  d550f666  c8c347a7  5a6ad9ad  43ada245  24e00850  f92939eb  78ce7989
t =  5  2b4209f5  04409a6a  d550f666  c8c347a7  714260ad  43ada245  24e00850  f92939eb
t =  6  e5030380  2b4209f5  04409a6a  d550f666  9b27a401  714260ad  43ada245  24e00850
t =  7  85a07b5f  e5030380  2b4209f5  04409a6a  0c657a79  9b27a401  714260ad  43ada245
t =  8  8e04ecb9  85a07b5f  e5030380  2b4209f5  32ca2d8c  0c657a79  9b27a401  714260ad
t =  9  8c87346b  8e04ecb9  85a07b5f  e5030380  1cc92596  32ca2d8c  0c657a79  9b27a401
t = 10  4798a3f4  8c87346b  8e04ecb9  85a07b5f  436b23e8  1cc92596  32ca2d8c  0c657a79
t = 11  f71fc5a9  4798a3f4  8c87346b  8e04ecb9  816fd6e9  436b23e8  1cc92596  32ca2d8c
t = 12  87912990  f71fc5a9  4798a3f4  8c87346b  1e578218  816fd6e9  436b23e8  1cc92596
t = 13  d932eb16  87912990  f71fc5a9  4798a3f4  745a48de  1e578218  816fd6e9  436b23e8
t = 14  c0645fde  d932eb16  87912990  f71fc5a9  0b92f20c  745a48de  1e578218  816fd6e9
t = 15  b0fa238e  c0645fde  d932eb16  87912990  07590dcd  0b92f20c  745a48de  1e578218
t = 16  21da9a9b  b0fa238e  c0645fde  d932eb16  8034229c  07590dcd  0b92f20c  745a48de
t = 17  c2fbd9d1  21da9a9b  b0fa238e  c0645fde  846ee454  8034229c  07590dcd  0b92f20c
t = 18  fe777bbf  c2fbd9d1  21da9a9b  b0fa238e  cc899961  846ee454  8034229c  07590dcd
t = 19  e1f20c33  fe777bbf  c2fbd9d1  21da9a9b  b0638179  cc899961  846ee454  8034229c
t = 20  9dc68b63  e1f20c33  fe777bbf  c2fbd9d1  8ada8930  b0638179  cc899961  846ee454
t = 21  c2606d6d  9dc68b63  e1f20c33  fe777bbf  e1257970  8ada8930  b0638179  cc899961
t = 22  a7a3623f  c2606d6d  9dc68b63  e1f20c33  49f5114a  e1257970  8ada8930  b0638179
t = 23  c5d53d8d  a7a3623f  c2606d6d  9dc68b63  aa47c347  49f5114a  e1257970  8ada8930
t = 24  1c2c2838  c5d53d8d  a7a3623f  c2606d6d  2823ef91  aa47c347  49f5114a  e1257970
t = 25  cde8037d  1c2c2838  c5d53d8d  a7a3623f  14383d8e  2823ef91  aa47c347  49f5114a
t = 26  b62ec4bc  cde8037d  1c2c2838  c5d53d8d  c74c6516  14383d8e  2823ef91  aa47c347
t = 27  77d37528  b62ec4bc  cde8037d  1c2c2838  edffbff8  c74c6516  14383d8e  2823ef91
t = 28  363482c9  77d37528  b62ec4bc  cde8037d  6112a3b7  edffbff8  c74c6516  14383d8e
t = 29  a0060b30  363482c9  77d37528  b62ec4bc  ade79437  6112a3b7  edffbff8  c74c6516
t = 30  ea992a22  a0060b30  363482c9  77d37528  0109ab3a  ade79437  6112a3b7  edffbff8
t = 31  73b33bf5  ea992a22  a0060b30  363482c9  ba591112  0109ab3a  ade79437  6112a3b7
t = 32  98e12507  73b33bf5  ea992a22  a0060b30  9cd9f5f6  ba591112  0109ab3a  ade79437
t = 33  fe604df5  98e12507  73b33bf5  ea992a22  59249dd3  9cd9f5f6  ba591112  0109ab3a
t = 34  a9a7738c  fe604df5  98e12507  73b33bf5  085f3833  59249dd3  9cd9f5f6  ba591112
t = 35  65a0cfe4  a9a7738c  fe604df5  98e12507  f4b002d6  085f3833  59249dd3  9cd9f5f6
t = 36  41a65cb1  65a0cfe4  a9a7738c  fe604df5  0772a26b  f4b002d6  085f3833  59249dd3
t = 37  34df1604  41a65cb1  65a0cfe4  a9a7738c  a507a53d  0772a26b  f4b002d6  085f3833
t = 38  6dc57a8a  34df1604  41a65cb1  65a0cfe4  f0781bc8  a507a53d  0772a26b  f4b002d6
t = 39  79ea687a  6dc57a8a  34df1604  41a65cb1  1efbc0a0  f0781bc8  a507a53d  0772a26b
t = 40  d6670766  79ea687a  6dc57a8a  34df1604  26352d63  1efbc0a0  f0781bc8  a507a53d
t = 41  df46652f  d6670766  79ea687a  6dc57a8a  838b2711  26352d63  1efbc0a0  f0781bc8
t = 42  17aa0dfe  df46652f  d6670766  79ea687a  decd4715  838b2711  26352d63  1efbc0a0
t = 43  9d4baf93  17aa0dfe  df46652f  d6670766  fda24c2e  decd4715  838b2711  26352d63
t = 44  26628815  9d4baf93  17aa0dfe  df46652f  a80f11f0  fda24c2e  decd4715  838b2711
t = 45  72ab4b91  26628815  9d4baf93  17aa0dfe  b7755da1  a80f11f0  fda24c2e  decd4715
t = 46  a14c14b0  72ab4b91  26628815  9d4baf93  d57b94a9  b7755da1  a80f11f0  fda24c2e
t = 47  4172328d  a14c14b0  72ab4b91  26628815  fecf0bc6  d57b94a9  b7755da1  a80f11f0
t = 48  05757ceb  4172328d  a14c14b0  72ab4b91  bd714038  fecf0bc6  d57b94a9  b7755da1
t = 49  f11bfaa8  05757ceb  4172328d  a14c14b0  6e5c390c  bd714038  fecf0bc6  d57b94a9
t = 50  7a0508a1  f11bfaa8  05757ceb  4172328d  52f1ccf7  6e5c390c  bd714038  fecf0bc6
t = 51  886e7a22  7a0508a1  f11bfaa8  05757ceb  49231c1e  52f1ccf7  6e5c390c  bd714038
t = 52  101fd28f  886e7a22  7a0508a1  f11bfaa8  529e7d00  49231c1e  52f1ccf7  6e5c390c
t = 53  f5702fdb  101fd28f  886e7a22  7a0508a1  9f4787c3  529e7d00  49231c1e  52f1ccf7
t = 54  3ec45cdb  f5702fdb  101fd28f  886e7a22  e50e1b4f  9f4787c3  529e7d00  49231c1e
t = 55  38cc9913  3ec45cdb  f5702fdb  101fd28f  54cb266b  e50e1b4f  9f4787c3  529e7d00
t = 56  fcd1887b  38cc9913  3ec45cdb  f5702fdb  9b5e906c  54cb266b  e50e1b4f  9f4787c3
t = 57  c062d46f  fcd1887b  38cc9913  3ec45cdb  7e44008e  9b5e906c  54cb266b  e50e1b4f
t = 58  ffb70472  c062d46f  fcd1887b  38cc9913  6d83bfc6  7e44008e  9b5e906c  54cb266b
t = 59  b6ae8fff  ffb70472  c062d46f  fcd1887b  b21bad3d  6d83bfc6  7e44008e  9b5e906c
t = 60  b85e2ce9  b6ae8fff  ffb70472  c062d46f  961f4894  b21bad3d  6d83bfc6  7e44008e
t = 61  04d24d6c  b85e2ce9  b6ae8fff  ffb70472  948d25b6  961f4894  b21bad3d  6d83bfc6
t = 62  d39a2165  04d24d6c  b85e2ce9  b6ae8fff  fb121210  948d25b6  961f4894  b21bad3d
t = 63  506e3058  d39a2165  04d24d6c  b85e2ce9  5ef50f24  fb121210  948d25b6  961f4894

Block 1 has been processed. The values of {Hi} are

H1 = 6a09e667 + 506e3058 = ba7816bf
H2 = bb67ae85 + d39a2165 = 8f01cfea
H3 = 3c6ef372 + 04d24d6c = 414140de
H4 = a54ff53a + b85e2ce9 = 5dae2223
H5 = 510e527f + 5ef50f24 = b00361a3
H6 = 9b05688c + fb121210 = 96177a9c
H7 = 1f83d9ab + 948d25b6 = b410ff61
H8 = 5be0cd19 + 961f4894 = f20015ad

The message digest is

    ba7816bf 8f01cfea 414140de 5dae2223 b00361a3 96177a9c b410ff61 f20015ad.
```

### Hash of "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq" (block 1)

```
            a         b         c         d         e         f         g         h

init:   6a09e667  bb67ae85  3c6ef372  a54ff53a  510e527f  9b05688c  1f83d9ab  5be0cd19
t =  0  5d6aebb1  6a09e667  bb67ae85  3c6ef372  fa2a4606  510e527f  9b05688c  1f83d9ab
t =  1  2f2d5fcf  5d6aebb1  6a09e667  bb67ae85  4eb1cfce  fa2a4606  510e527f  9b05688c
t =  2  97651825  2f2d5fcf  5d6aebb1  6a09e667  62d5c49e  4eb1cfce  fa2a4606  510e527f
t =  3  4a8d64d5  97651825  2f2d5fcf  5d6aebb1  6494841b  62d5c49e  4eb1cfce  fa2a4606
t =  4  f921c212  4a8d64d5  97651825  2f2d5fcf  05c4f88a  6494841b  62d5c49e  4eb1cfce
t =  5  55c8ef48  f921c212  4a8d64d5  97651825  7ff91c94  05c4f88a  6494841b  62d5c49e
t =  6  485835b7  55c8ef48  f921c212  4a8d64d5  39a5b2ca  7ff91c94  05c4f88a  6494841b
t =  7  d237e6db  485835b7  55c8ef48  f921c212  a401d211  39a5b2ca  7ff91c94  05c4f88a
t =  8  359f2bce  d237e6db  485835b7  55c8ef48  c09ffec4  a401d211  39a5b2ca  7ff91c94
t =  9  3a474b2b  359f2bce  d237e6db  485835b7  9037b3b8  c09ffec4  a401d211  39a5b2ca
t = 10  b8e2b4cb  3a474b2b  359f2bce  d237e6db  443ed29e  9037b3b8  c09ffec4  a401d211
t = 11  1762215c  b8e2b4cb  3a474b2b  359f2bce  ee1c97a8  443ed29e  9037b3b8  c09ffec4
t = 12  101a4861  1762215c  b8e2b4cb  3a474b2b  839a0fc9  ee1c97a8  443ed29e  9037b3b8
t = 13  d68e6457  101a4861  1762215c  b8e2b4cb  9243f8af  839a0fc9  ee1c97a8  443ed29e
t = 14  dd16cbb3  d68e6457  101a4861  1762215c  9162aded  9243f8af  839a0fc9  ee1c97a8
t = 15  c3486194  dd16cbb3  d68e6457  101a4861  1496a54f  9162aded  9243f8af  839a0fc9
t = 16  b9dcacb1  c3486194  dd16cbb3  d68e6457  d4f64250  1496a54f  9162aded  9243f8af
t = 17  046a193e  b9dcacb1  c3486194  dd16cbb3  885370b6  d4f64250  1496a54f  9162aded
t = 18  f402f058  046a193e  b9dcacb1  c3486194  6f433549  885370b6  d4f64250  1496a54f
t = 19  2139187b  f402f058  046a193e  b9dcacb1  7c304206  6f433549  885370b6  d4f64250
t = 20  d70ac17d  2139187b  f402f058  046a193e  7cc6b262  7c304206  6f433549  885370b6
t = 21  1b2b66b8  d70ac17d  2139187b  f402f058  d560b028  7cc6b262  7c304206  6f433549
t = 22  ae2e2d4f  1b2b66b8  d70ac17d  2139187b  f074fc95  d560b028  7cc6b262  7c304206
t = 23  59fce6b9  ae2e2d4f  1b2b66b8  d70ac17d  a2c7d51d  f074fc95  d560b028  7cc6b262
t = 24  4a885065  59fce6b9  ae2e2d4f  1b2b66b8  763597fb  a2c7d51d  f074fc95  d560b028
t = 25  573221da  4a885065  59fce6b9  ae2e2d4f  36e74eb4  763597fb  a2c7d51d  f074fc95
t = 26  128661da  573221da  4a885065  59fce6b9  1162d575  36e74eb4  763597fb  a2c7d51d
t = 27  73f858af  128661da  573221da  4a885065  e77c797f  1162d575  36e74eb4  763597fb
t = 28  74bcf468  73f858af  128661da  573221da  72abaecd  e77c797f  1162d575  36e74eb4
t = 29  df7151a0  74bcf468  73f858af  128661da  7629c961  72abaecd  e77c797f  1162d575
t = 30  eb43f3ed  df7151a0  74bcf468  73f858af  0635d880  7629c961  72abaecd  e77c797f
t = 31  5581ab07  eb43f3ed  df7151a0  74bcf468  df980085  0635d880  7629c961  72abaecd
t = 32  9fc905c8  5581ab07  eb43f3ed  df7151a0  a94d2af1  df980085  0635d880  7629c961
t = 33  9ce5a62f  9fc905c8  5581ab07  eb43f3ed  6ef3b6bd  a94d2af1  df980085  0635d880
t = 34  1df8e885  9ce5a62f  9fc905c8  5581ab07  2a9e048e  6ef3b6bd  a94d2af1  df980085
t = 35  0786dce8  1df8e885  9ce5a62f  9fc905c8  de2a21d1  2a9e048e  6ef3b6bd  a94d2af1
t = 36  2c55d3a6  0786dce8  1df8e885  9ce5a62f  b067c1af  de2a21d1  2a9e048e  6ef3b6bd
t = 37  a985b4be  2c55d3a6  0786dce8  1df8e885  f72bf353  b067c1af  de2a21d1  2a9e048e
t = 38  91ac9d5d  a985b4be  2c55d3a6  0786dce8  68d8d590  f72bf353  b067c1af  de2a21d1
t = 39  7e4d30b8  91ac9d5d  a985b4be  2c55d3a6  9f5b9b6d  68d8d590  f72bf353  b067c1af
t = 40  7e056794  7e4d30b8  91ac9d5d  a985b4be  423b26c0  9f5b9b6d  68d8d590  f72bf353
t = 41  508a16ab  7e056794  7e4d30b8  91ac9d5d  45459d97  423b26c0  9f5b9b6d  68d8d590
t = 42  b62c7013  508a16ab  7e056794  7e4d30b8  80a92a00  45459d97  423b26c0  9f5b9b6d
t = 43  167361de  b62c7013  508a16ab  7e056794  41dd3844  80a92a00  45459d97  423b26c0
t = 44  de71e2f2  167361de  b62c7013  508a16ab  ff61c636  41dd3844  80a92a00  45459d97
t = 45  18f0d19d  de71e2f2  167361de  b62c7013  6b88472c  ff61c636  41dd3844  80a92a00
t = 46  165be9cd  18f0d19d  de71e2f2  167361de  a483f080  6b88472c  ff61c636  41dd3844
t = 47  13d82741  165be9cd  18f0d19d  de71e2f2  a7802a4d  a483f080  6b88472c  ff61c636
t = 48  017b9d99  13d82741  165be9cd  18f0d19d  aeb10b60  a7802a4d  a483f080  6b88472c
t = 49  543c99a1  017b9d99  13d82741  165be9cd  16f134b6  aeb10b60  a7802a4d  a483f080
t = 50  758ca97a  543c99a1  017b9d99  13d82741  100cf2ea  16f134b6  aeb10b60  a7802a4d
t = 51  81c1cde0  758ca97a  543c99a1  017b9d99  5c47eb7b  100cf2ea  16f134b6  aeb10b60
t = 52  b8d55619  81c1cde0  758ca97a  543c99a1  1c806a61  5c47eb7b  100cf2ea  16f134b6
t = 53  1d6de87a  b8d55619  81c1cde0  758ca97a  3443bed4  1c806a61  5c47eb7b  100cf2ea
t = 54  f907b313  1d6de87a  b8d55619  81c1cde0  61a41711  3443bed4  1c806a61  5c47eb7b
t = 55  9e57c4a0  f907b313  1d6de87a  b8d55619  eec13548  61a41711  3443bed4  1c806a61
t = 56  71629856  9e57c4a0  f907b313  1d6de87a  2f6c8c4e  eec13548  61a41711  3443bed4
t = 57  7c015a2c  71629856  9e57c4a0  f907b313  cb9d3dd0  2f6c8c4e  eec13548  61a41711
t = 58  921fccb6  7c015a2c  71629856  9e57c4a0  43d8a034  cb9d3dd0  2f6c8c4e  eec13548
t = 59  e18f259a  921fccb6  7c015a2c  71629856  51e15869  43d8a034  cb9d3dd0  2f6c8c4e
t = 60  bcfce922  e18f259a  921fccb6  7c015a2c  962d8621  51e15869  43d8a034  cb9d3dd0
t = 61  f6f443f8  bcfce922  e18f259a  921fccb6  acc75916  962d8621  51e15869  43d8a034
t = 62  86126910  f6f443f8  bcfce922  e18f259a  2fc08f85  acc75916  962d8621  51e15869
t = 63  1bdc6f6f  86126910  f6f443f8  bcfce922  25d2430a  2fc08f85  acc75916  962d8621

Block 1 has been processed. The values of {Hi} are

H1 = 6a09e667 + 1bdc6f6f = 85e655d6
H2 = bb67ae85 + 86126910 = 417a1795
H3 = 3c6ef372 + f6f443f8 = 3363376a
H4 = a54ff53a + bcfce922 = 624cde5c
H5 = 510e527f + 25d2430a = 76e09589
H6 = 9b05688c + 2fc08f85 = cac5f811
H7 = 1f83d9ab + acc75916 = cc4b32c1
H8 = 5be0cd19 + 962d8621 = f20e533a
```

### Hash of "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq" (block 2)

```
            a         b         c         d         e         f         g         h

init:   85e655d6  417a1795  3363376a  624cde5c  76e09589  cac5f811  cc4b32c1  f20e533a
t =  0  7c20c838  85e655d6  417a1795  3363376a  4670ae6e  76e09589  cac5f811  cc4b32c1
t =  1  7c3c0f86  7c20c838  85e655d6  417a1795  8c51be64  4670ae6e  76e09589  cac5f811
t =  2  fd1eebdc  7c3c0f86  7c20c838  85e655d6  af71b9ea  8c51be64  4670ae6e  76e09589
t =  3  f268faa9  fd1eebdc  7c3c0f86  7c20c838  e20362ef  af71b9ea  8c51be64  4670ae6e
t =  4  185a5d79  f268faa9  fd1eebdc  7c3c0f86  8dff3001  e20362ef  af71b9ea  8c51be64
t =  5  3eeb6c06  185a5d79  f268faa9  fd1eebdc  fe20cda6  8dff3001  e20362ef  af71b9ea
t =  6  89bba3f1  3eeb6c06  185a5d79  f268faa9  0a34df03  fe20cda6  8dff3001  e20362ef
t =  7  bf9a93a0  89bba3f1  3eeb6c06  185a5d79  059abdd1  0a34df03  fe20cda6  8dff3001
t =  8  2c096744  bf9a93a0  89bba3f1  3eeb6c06  abfa465b  059abdd1  0a34df03  fe20cda6
t =  9  2d964e86  2c096744  bf9a93a0  89bba3f1  aa27ed82  abfa465b  059abdd1  0a34df03
t = 10  5b35025b  2d964e86  2c096744  bf9a93a0  10e77723  aa27ed82  abfa465b  059abdd1
t = 11  5eb4ec40  5b35025b  2d964e86  2c096744  e11b4548  10e77723  aa27ed82  abfa465b
t = 12  35ee996d  5eb4ec40  5b35025b  2d964e86  5c24e2a2  e11b4548  10e77723  aa27ed82
t = 13  d74080fa  35ee996d  5eb4ec40  5b35025b  68aa893f  5c24e2a2  e11b4548  10e77723
t = 14  0cea5cbc  d74080fa  35ee996d  5eb4ec40  60356548  68aa893f  5c24e2a2  e11b4548
t = 15  16a8cc79  0cea5cbc  d74080fa  35ee996d  0fcb1f6f  60356548  68aa893f  5c24e2a2
t = 16  f16f634e  16a8cc79  0cea5cbc  d74080fa  8b21cdc1  0fcb1f6f  60356548  68aa893f
t = 17  23dcb6c2  f16f634e  16a8cc79  0cea5cbc  ca9182d3  8b21cdc1  0fcb1f6f  60356548
t = 18  dcff40fd  23dcb6c2  f16f634e  16a8cc79  69bf7b95  ca9182d3  8b21cdc1  0fcb1f6f
t = 19  76f1a2bc  dcff40fd  23dcb6c2  f16f634e  0dc84bb1  69bf7b95  ca9182d3  8b21cdc1
t = 20  20aad899  76f1a2bc  dcff40fd  23dcb6c2  cc4769f2  0dc84bb1  69bf7b95  ca9182d3
t = 21  d44dc81a  20aad899  76f1a2bc  dcff40fd  5bace62d  cc4769f2  0dc84bb1  69bf7b95
t = 22  f13ae55b  d44dc81a  20aad899  76f1a2bc  966aa287  5bace62d  cc4769f2  0dc84bb1
t = 23  a4195b91  f13ae55b  d44dc81a  20aad899  eddbd6ed  966aa287  5bace62d  cc4769f2
t = 24  4984fa79  a4195b91  f13ae55b  d44dc81a  a530d939  eddbd6ed  966aa287  5bace62d
t = 25  aa6cb982  4984fa79  a4195b91  f13ae55b  0b5eeea4  a530d939  eddbd6ed  966aa287
t = 26  9450fbbc  aa6cb982  4984fa79  a4195b91  09166dda  0b5eeea4  a530d939  eddbd6ed
t = 27  0d936bab  9450fbbc  aa6cb982  4984fa79  6e495d4b  09166dda  0b5eeea4  a530d939
t = 28  d958b529  0d936bab  9450fbbc  aa6cb982  c2fa99b1  6e495d4b  09166dda  0b5eeea4
t = 29  1cfa5eb0  d958b529  0d936bab  9450fbbc  6c49db9f  c2fa99b1  6e495d4b  09166dda
t = 30  02ef3a5f  1cfa5eb0  d958b529  0d936bab  5da10665  6c49db9f  c2fa99b1  6e495d4b
t = 31  b0eab1c5  02ef3a5f  1cfa5eb0  d958b529  f6d93952  5da10665  6c49db9f  c2fa99b1
t = 32  0bfba73c  b0eab1c5  02ef3a5f  1cfa5eb0  8b99e3a9  f6d93952  5da10665  6c49db9f
t = 33  4bd1df96  0bfba73c  b0eab1c5  02ef3a5f  905e44ac  8b99e3a9  f6d93952  5da10665
t = 34  9907f1b6  4bd1df96  0bfba73c  b0eab1c5  66c3043d  905e44ac  8b99e3a9  f6d93952
t = 35  ecde4e0d  9907f1b6  4bd1df96  0bfba73c  5dc119e6  66c3043d  905e44ac  8b99e3a9
t = 36  2f11c939  ecde4e0d  9907f1b6  4bd1df96  fed4ce1d  5dc119e6  66c3043d  905e44ac
t = 37  d949682b  2f11c939  ecde4e0d  9907f1b6  32d99008  fed4ce1d  5dc119e6  66c3043d
t = 38  adca7a96  d949682b  2f11c939  ecde4e0d  c6cce4ff  32d99008  fed4ce1d  5dc119e6
t = 39  221b8a5a  adca7a96  d949682b  2f11c939  0b82c5eb  c6cce4ff  32d99008  fed4ce1d
t = 40  12d97845  221b8a5a  adca7a96  d949682b  e4213ca2  0b82c5eb  c6cce4ff  32d99008
t = 41  2c794876  12d97845  221b8a5a  adca7a96  ff6759ba  e4213ca2  0b82c5eb  c6cce4ff
t = 42  8300fca2  2c794876  12d97845  221b8a5a  e0e3457c  ff6759ba  e4213ca2  0b82c5eb
t = 43  f2ad6322  8300fca2  2c794876  12d97845  cc48c7f3  e0e3457c  ff6759ba  e4213ca2
t = 44  0f154e11  f2ad6322  8300fca2  2c794876  6f9517cb  cc48c7f3  e0e3457c  ff6759ba
t = 45  104a7db4  0f154e11  f2ad6322  8300fca2  5348e8f6  6f9517cb  cc48c7f3  e0e3457c
t = 46  0b3303a7  104a7db4  0f154e11  f2ad6322  bbe1c39a  5348e8f6  6f9517cb  cc48c7f3
t = 47  d7354d5b  0b3303a7  104a7db4  0f154e11  aad55b6b  bbe1c39a  5348e8f6  6f9517cb
t = 48  b736d7a6  d7354d5b  0b3303a7  104a7db4  68f25260  aad55b6b  bbe1c39a  5348e8f6
t = 49  2748e5ec  b736d7a6  d7354d5b  0b3303a7  d4b58576  68f25260  aad55b6b  bbe1c39a
t = 50  d8aabcf9  2748e5ec  b736d7a6  d7354d5b  27844711  d4b58576  68f25260  aad55b6b
t = 51  1a6bcf6a  d8aabcf9  2748e5ec  b736d7a6  ff5e99d0  27844711  d4b58576  68f25260
t = 52  4eca6fa0  1a6bcf6a  d8aabcf9  2748e5ec  989ed071  ff5e99d0  27844711  d4b58576
t = 53  ec02560a  4eca6fa0  1a6bcf6a  d8aabcf9  7151df8e  989ed071  ff5e99d0  27844711
t = 54  d9f0c115  ec02560a  4eca6fa0  1a6bcf6a  624150c4  7151df8e  989ed071  ff5e99d0
t = 55  92952710  d9f0c115  ec02560a  4eca6fa0  226806d6  624150c4  7151df8e  989ed071
t = 56  20d4d0e4  92952710  d9f0c115  ec02560a  4e515a4d  226806d6  624150c4  7151df8e
t = 57  4348eb1f  20d4d0e4  92952710  d9f0c115  c21eddf9  4e515a4d  226806d6  624150c4
t = 58  286fe5f0  4348eb1f  20d4d0e4  92952710  54076664  c21eddf9  4e515a4d  226806d6
t = 59  1c4cddd9  286fe5f0  4348eb1f  20d4d0e4  f487a853  54076664  c21eddf9  4e515a4d
t = 60  a9f181dd  1c4cddd9  286fe5f0  4348eb1f  27ccb387  f487a853  54076664  c21eddf9
t = 61  b25cef29  a9f181dd  1c4cddd9  286fe5f0  2aa1bb13  27ccb387  f487a853  54076664
t = 62  908c2123  b25cef29  a9f181dd  1c4cddd9  9a392956  2aa1bb13  27ccb387  f487a853
t = 63  9ea7148b  908c2123  b25cef29  a9f181dd  2c5c4ed0  9a392956  2aa1bb13  27ccb387

Block 2 has been processed. The values of {Hi} are

H1 = 85e655d6 + 9ea7148b = 248d6a61
H2 = 417a1795 + 908c2123 = d20638b8
H3 = 3363376a + b25cef29 = e5c02693
H4 = 624cde5c + a9f181dd = 0c3e6039
H5 = 76e09589 + 2c5c4ed0 = a33ce459
H6 = cac5f811 + 9a392956 = 64ff2167
H7 = cc4b32c1 + 2aa1bb13 = f6ecedd4
H8 = f20e533a + 27ccb387 = 19db06c1

The message digest is

    248d6a61 d20638b8 e5c02693 0c3e6039 a33ce459 64ff2167 f6ecedd4 19db06c1.
```

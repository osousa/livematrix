let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/Desktop/Go/gobotrix/livematrix/livematrix/serverbot
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
set shortmess=aoO
argglobal
%argdel
$argadd go.mod
set stal=2
tabnew
tabnew
tabnew
tabnew
tabnew
tabrewind
edit main.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
enew
file NvimTree_1
balt main.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
wincmd w
argglobal
balt main.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 9 - ((8 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 9
normal! 0
wincmd w
2wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext
edit NvimTree_14
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
balt chat/session.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
wincmd w
argglobal
if bufexists("chat/session.go") | buffer chat/session.go | else | edit chat/session.go | endif
if &buftype ==# 'terminal'
  silent file chat/session.go
endif
balt chat/database.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 57 - ((52 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 57
normal! 05|
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext
edit NvimTree_15
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
balt chat/server.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
wincmd w
argglobal
if bufexists("chat/server.go") | buffer chat/server.go | else | edit chat/server.go | endif
if &buftype ==# 'terminal'
  silent file chat/server.go
endif
balt chat/client.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 139 - ((52 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 139
normal! 013|
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext
edit NvimTree_21
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
balt chat/database.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
wincmd w
argglobal
if bufexists("chat/database.go") | buffer chat/database.go | else | edit chat/database.go | endif
if &buftype ==# 'terminal'
  silent file chat/database.go
endif
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 129 - ((3 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 129
normal! 0
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext
edit NvimTree_17
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
balt chat/client.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
wincmd w
argglobal
if bufexists("chat/client.go") | buffer chat/client.go | else | edit chat/client.go | endif
if &buftype ==# 'terminal'
  silent file chat/client.go
endif
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 71 - ((32 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 71
normal! 038|
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext
edit NvimTree_18
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
argglobal
balt chat/message.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
wincmd w
argglobal
if bufexists("chat/message.go") | buffer chat/message.go | else | edit chat/message.go | endif
if &buftype ==# 'terminal'
  silent file chat/message.go
endif
balt chat/message.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 9 - ((8 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 9
normal! 0
wincmd w
exe 'vert 1resize ' . ((&columns * 30 + 119) / 238)
exe 'vert 2resize ' . ((&columns * 207 + 119) / 238)
tabnext 1
set stal=1
badd +9 main.go
badd +0 go.mod
badd +91 chat/session.go
badd +1 chat/server.go
badd +10 chat/message.go
badd +0 term://~/Desktop/Go/gobotrix/servebot//222026:/bin/bash
badd +0 NvimTree_14
badd +188 chat/database.go
badd +0 NvimTree_15
badd +76 chat/client.go
badd +0 NvimTree_21
badd +0 NvimTree_17
badd +0 NvimTree_18
badd +1 .env
badd +1 models.go
badd +141 database.go
badd +26 term://~/Desktop/Go/gobotrix/servebot//604892:/bin/bash
badd +1 /usr/local/go/src/builtin/builtin.go
badd +26 term://~/Desktop/Go/gobotrix/servebot//402751:/bin/bash
badd +4 term://~/Desktop/Go/gobotrix/servebot//1170943:/bin/bash
badd +37 term://~/Desktop/Go/gobotrix/servebot//400562:/bin/bash
badd +41 term://~/Desktop/Go/gobotrix/servebot//401069:/bin/bash
badd +83 term://~/Desktop/Go/gobotrix/servebot//404111:/bin/bash
badd +2 chat/go.mod
badd +2 term://~/Desktop/Go/gobotrix/servebot//730695:/bin/bash
badd +28 start.sh
badd +197 term://~/Desktop/Go/gobotrix/servebot//85499:/bin/bash
badd +9 term://~/Desktop/Go/gobotrix/servebot//217743:/bin/bash
badd +31 term://~/Desktop/Go/gobotrix/servebot//219513:/bin/bash
badd +2 chat/models.go
badd +1 session.go
badd +2 term://~/Desktop/Go/gobotrix/servebot//250780:/bin/bash
badd +56 term://~/Desktop/Go/gobotrix/servebot//265116:/bin/bash
badd +51 term://~/Desktop/Go/gobotrix/servebot//291595:/bin/bash
badd +48 term://~/Desktop/Go/gobotrix/servebot//301879:/bin/bash
badd +37 ~/.vimrc
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20 shortmess=filnxtToOF
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
nohlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :

syntax on
"set backspace = index,eol,start
"set listchars=extends:>,precedes:<
"set listchars=tab:>-,extends:>,precedes:<
"eol:$
"set list
set nu

nnoremap zz :update<CR>
set tabstop=2 expandtab
set shiftwidth=2
inoremap jj <esc>

noremap <F11> :set list!<CR>


"nnoremap yy yy"+yy
"vnoremap y ygv"+y


" Copy to clipboard
"vnoremap  <leader>y  "+y
"nnoremap  <leader>y  "+y

" Paste from clipboard
"nnoremap <leader>p "+p
"vnoremap <leader>p "+p

autocmd TextYankPost * if v:event.operator ==# 'y' | call system('/mnt/c/Windows/System32/clip.exe', @0) | endif

set mouse=a
colorscheme morning

set belloff=all
set autoindent


set encoding=utf-8
set fileencodings=utf-8,cp949


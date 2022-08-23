# Resume Materi KMGolang – Version Control and Branch Management (Git)

Tuliskan 3 poin yang dipelajari dari materi tersebut. Resume / ringkasan materi dapat disubmit melalui Github

## 1. Pentingnya Version Control System

Version Control System (VCS) memperbolehkan kita untuk mengembalikan berkas-berkas ke keadaan sebelumnya, mengembalikan seluruh proyek kembali ke keadaan sebelumnya, membandingkan perubahan-perubahan di setiap waktu, melihat siapa yang terakhir mengubah sesuatu yang mungkin menimbulkan masalah, siapa dan kapan yang mengenalkan sebuah isu dan banyak lagi.

## 2. Git

Git merupakan salah satu VCS populer pada proyek perangkat lunak yang diciptakan oleh Linus Torvalds (2005). Sistem yang dianut adalah sistem terdistribusi (bukan tersentralisasi) artinya semua harus sinkron kepada remote server. Salah satu git hosting service populer adalah Github.

### Perintah/command pada VCS Git

Perintah-perintah dasar pada git adalah sebagai berikut

| Command                                           | Fungsi                                               |
|---------------------------------------------------|------------------------------------------------------|
| [git config](https://git-scm.com/docs/git-config) | Mengatur konfigurasi tertentu, seperti `nama, email` |
| [git init](https://git-scm.com/docs/git-init)     | Melakukan membuat repository baru                    |
| [git add](https://git-scm.com/docs/git-add)       | Menambahkan file ke index                            |
| [git clone](https://git-scm.com/docs/git-clone)   | Membuat salinan repository                           |
| [git remote](https://git-scm.com/docs/git-remote) | Perintah berkaitan mengelola remote repository       |
| [git push](https://git-scm.com/docs/git-push) | Mengupdate project yang ada di git server |
| [git status](https://git-scm.com/docs/git-status) | Melihat status dari git |
| [git commit](https://git-scm.com/docs/git-commit) | Menyimpan perubahan ke repository |
| [git diff](https://git-scm.com/docs/git-diff) | Melihat perubahan diantara commit |
| [git stash](https://git-scm.com/docs/git-stash) | Menyimpan perubahan di stash directory |
| [git log](https://git-scm.com/docs/git-log) | Menampilkan commit logs |
| [git checkout](https://git-scm.com/docs/git-checkout) | Mengubah branch atau mengembalikan file working tree |
| [git reset](https://git-scm.com/docs/git-reset) | Reset HEAD ke kondisi tertentu |
| [git fetch](https://git-scm.com/docs/git-fetch) | Download objek dan refs dari repository lainnya |
| [git pull](https://git-scm.com/docs/git-pull) | Mengambil dari repository dan diintegrasikan |
| [git branch](https://git-scm.com/docs/git-branch) | Membuat, menghapus, mendata cabang (branches) |
| [git merge](https://git-scm.com/docs/git-merge) | Menggabungkan 2 atau lebih riwayat development |

## 3. Workflow Collaboration

Workflow collaboration bertujuan untuk memahami alur kerja yang tepat untuk berkolaborasi dengan GitHub atau Gitlab.

Cara mengoptimalkan kolaborasi dalam development:

1. Buat branch `MASTER` dari branch development
2. Hindari edit langsung pada saat develop
3. Merge branch `FEATURE` hanya ke branch `DEVELOPMENT`
4. Merge branch `DEVELOPMENT` ke branch `MASTER` jika sudah selesai
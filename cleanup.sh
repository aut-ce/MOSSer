#!/bin/bash
# In The Name Of God
# ========================================
# [] File Name : cleanup.sh
#
# [] Creation Date : 28-06-2018
#
# [] Created By : Parham Alvani <parham.alvani@gmail.com>
# =======================================
ls -a | grep -v -E 'moss.pl|README.md|run.sh|.gitignore|cleanup.sh|.git' | xargs rm -Rf

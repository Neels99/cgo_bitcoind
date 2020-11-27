printf "Start build golang project...\n"
bash c_archive_build.sh
printf "Start build c++ project...\n"
bash cpp_proj_build.sh
printf "Run project:\n"
printf "_____________________________\n"
./main
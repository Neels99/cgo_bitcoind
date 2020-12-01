cmake_minimum_required(VERSION 2.8)	 # Проверка версии CMake.
										# Если версия установленой программы
										# старее указаной, произайдёт аварийный выход.

project(hello_world)			# Название проекта

set(SOURCE_EXE Test_C_GBCI.cpp)		# Установка переменной со списком исходников для исполняемого файла

# find_package(CURL)

# if(WIN32)
#     add_definitions(-DCURL_STATICLIB)
#     link_directories("${CMAKE_SOURCE_DIR}/thirdparty/openssl/" "${CMAKE_SOURCE_DIR}/thirdparty/"curl_win32)
# else()
#     link_directories("${CMAKE_SOURCE_DIR}/thirdparty/curl2")
# endif()

set(CURL_LIBRARY "-lcurl") 
find_package(CURL REQUIRED) 

add_executable(main ${SOURCE_EXE})	# Создает исполняемый файл с именем main

include_directories(${CURL_INCLUDE_DIR})
target_link_libraries(main ${CURL_LIBRARIES})

#--------
# Look for required libraries
# SET(requiredlibs)

# FIND_PACKAGE(CURL)
# IF(CURL_FOUND)
#   INCLUDE_DIRECTORIES(${CURL_INCLUDE_DIR})
#   SET(requiredlibs ${requiredlibs} ${CURL_LIBRARIES} )
# ELSE(CURL_FOUND)
#   MESSAGE(FATAL_ERROR "Could not find the CURL library and development files.")
# ENDIF(CURL_FOUND)
#---------------------------------

# if(WIN32)
#     target_link_libraries(main openssl curl_win32)
# else()
#     target_link_libraries(main curl_win32)
# endif()

#target_link_libraries(main PRIVATE CURL::libcurl)
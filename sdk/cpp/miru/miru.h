// sdk/cpp/include/miru.h
#ifndef MIRU_H
#define MIRU_H

#include <string>

// Client class
class Client {
public:
    static Client from_cert(const std::string& ca_file, const std::string& cert_file, const std::string& key_file);

    template <typename T>
    Config<T> init_config(const std::string& slug, const std::string& schema_file, T data, bool development, const std::string& development_file);
};

// Config class
template <typename T>
class Config {
public:
    T data();
    void set(const T& data);
};

#endif // MIRU_H
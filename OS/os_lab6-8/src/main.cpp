#include <iostream>
#include <unistd.h>
#include <string>
#include <vector>
#include <set>
#include <sstream>
#include <signal.h>
#include "zmq.hpp"
#include "tree.h"

const int WAIT_TIME = 1000;
const int PORT_BASE  = 5050;

bool send_message(zmq::socket_t &socket, const std::string &message_string)
{
    zmq::message_t message(message_string.size());
    memcpy(message.data(), message_string.c_str(), message_string.size());
    return socket.send(message);
}

std::string recieve_message(zmq::socket_t &socket)
{
    zmq::message_t message;
    bool ok = false;
    try
    {
        ok = socket.recv(&message);
    }
    catch (...)
    {
        ok = false;
    }
    std::string recieved_message(static_cast<char*>(message.data()), message.size());
    if (recieved_message.empty() || !ok)
    {
        return "Error: Node is not available";
    }
    return recieved_message;
}

void create_node(int id, int port)
{
    char* arg0 = strdup("./client");
    char* arg1 = strdup((std::to_string(id)).c_str());
    char* arg2 = strdup((std::to_string(port)).c_str());
    char* args[] = {arg0, arg1, arg2, NULL};
    execv("./client", args);
}

std::string get_port_name(const int port)
{
    return "tcp://127.0.0.1:" + std::to_string(port);
}

bool is_number(std::string val)
{
    try
    {
        int tmp = std::stoi(val);
        return true;
    }
    catch(std::exception& e)
    {
        std::cout << "Error: " << e.what() << "\n";
        return false;
    }
}

int main()
{
    Tree T;
    std::string command;
    int child_pid = 0;
    int child_id = 0;
    zmq::context_t context(1);
    zmq::socket_t main_socket(context, ZMQ_REQ);
    std::cout << "Commands:\n";
    std::cout << "create id\n";
    std::cout << "exec id subcommand (start/stop/time)\n";
    std::cout << "kill id\n";
    std::cout << "pingall\n";
    std::cout << "exit\n" << std::endl;
    while(1)
    {
        std::cin >> command;
        if (command == "create")
        {
            size_t node_id = 0;
            std::string str = "";
            std::string result = "";
            std::cin >> str;
            if (!is_number(str))
            {
                continue;
            }
            node_id = stoi(str);
            if (child_pid == 0)
            {
                main_socket.bind(get_port_name(PORT_BASE + node_id));
                child_pid = fork();
                if (child_pid == -1)
                {
                    std::cout << "Unable to create first worker node\n";
                    child_pid = 0;
                    exit(1);
                } 
                else if (child_pid == 0)
                {
                    create_node(node_id, PORT_BASE + node_id);
                }
                else
                {
                    child_id = node_id;
                    send_message(main_socket,"pid");
                    result = recieve_message(main_socket);
                }
            }
            else
            {
                std::string msg_s = "create " + std::to_string(node_id);
                send_message(main_socket, msg_s);
                result = recieve_message(main_socket);
            }
            if (result.substr(0, 2) == "Ok")
            {
                T.push(node_id);
            }
            std::cout << result << "\n";
        } 
        else if (command == "kill")
        {
            int node_id = 0;
            std::string str = "";
            std::cin >> str;
            if (!is_number(str))
            {
                continue;
            }
            node_id = stoi(str);
            if (child_pid == 0)
            {
                std::cout << "Error: Not found\n";
                continue;
            }
            if (node_id == child_id)
            {
                kill(child_pid, SIGTERM);
                kill(child_pid, SIGKILL);
                child_id = 0;
                child_pid = 0;
                T.kill(node_id);
                std::cout << "Ok\n";
                continue;
            }
            std::string message_string = "kill " + std::to_string(node_id);
            send_message(main_socket, message_string);
            std::string recieved_message = recieve_message(main_socket);
            if (recieved_message.substr(0, std::min<int>(recieved_message.size(), 2)) == "Ok")
            {
                T.kill(node_id);
            }
            std::cout << recieved_message << "\n";
        }
        else if (command == "exec")
        {
            std::string id_str = "";
            std::string subcommand = "";
            int id = 0; 
            std::cin >> id_str >> subcommand;
            if (!is_number(id_str))
            {
                continue;
            }
            id = stoi(id_str);
            if ((subcommand != "start") && (subcommand != "stop") && (subcommand != "time"))
            {
                std::cout << "Wrong subcommandmand\n";
                continue;
            }
            std::string message_string = "exec " + std::to_string(id) + " " + subcommand;
            send_message(main_socket, message_string);
            std::string recieved_message = recieve_message(main_socket);
            std::cout << recieved_message << "\n";
        }
        else if (command == "pingall")
        {
            send_message(main_socket,"pingall");
            std::string recieved = recieve_message(main_socket);
            std::istringstream is;        
            if (recieved.substr(0, std::min<int>(recieved.size(), 5)) == "Error")
            {
                is = std::istringstream("");
            }
            else
            {
                is = std::istringstream(recieved);
            }        
            std::set<int> recieved_T;
            int rec_id;
            while (is >> rec_id)
            {
                recieved_T.insert(rec_id);
            }
            std::vector<int> from_tree = T.get_nodes();
            auto part_it = partition(from_tree.begin(), from_tree.end(), [&recieved_T] (int a)
            {
                return recieved_T.count(a) == 0;
            });
            if (part_it == from_tree.begin())
            {
                std::cout << "Ok:-1\n";
            }
            else
            {
                std::cout << "Ok:";
                for (auto it = from_tree.begin(); it != part_it; ++it)
                {
                    std::cout  << *it << " ";
                }
                std::cout << "\n";
            }   
        }
        else if (command == "exit")
        {
            int n = system("killall client");
            break; 
        }
    }
    return 0;
}

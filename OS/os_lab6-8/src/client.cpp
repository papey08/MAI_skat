#include <iostream>
#include <unistd.h>
#include <string>
#include <sstream>
#include <exception>
#include <signal.h>
#include "zmq.hpp"
#include "timer.h"

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

void rl_create(zmq::socket_t& parent_socket, zmq::socket_t& socket, int& create_id, int& id, int& pid)
{
    if (pid == -1) 
    {
        send_message(parent_socket, "Error: Cannot fork");
        pid = 0;
    } 
    else if (pid == 0) 
    {
        create_node(create_id,PORT_BASE + create_id);
    } 
    else
    {
        id = create_id;
        send_message(socket, "pid");
        send_message(parent_socket, recieve_message(socket));
    }
}

void rl_kill(zmq::socket_t& parent_socket, zmq::socket_t& socket,  int& delete_id, int& id, int& pid, std::string& request_string)
{
    if (id == 0)
    {
        send_message(parent_socket, "Error: Not found");
    } 
    else if (id == delete_id)
    {
        send_message(socket, "kill_children");
        recieve_message(socket);
        kill(pid,SIGTERM);
        kill(pid,SIGKILL);
        id = 0;
        pid = 0;
        send_message(parent_socket, "Ok");
    } 
    else 
    {
        send_message(socket, request_string);
        send_message(parent_socket, recieve_message(socket));
    }
}

void rl_exec(zmq::socket_t& parent_socket, zmq::socket_t& socket,  int& id, int& pid, std::string& request_string)
{
    if (pid == 0)
    {
        std::string recieve_message = "Error:" + std::to_string(id);
        recieve_message += ": Not found";
        send_message(parent_socket, recieve_message);
    } 
    else
    {
        send_message(socket, request_string);
        send_message(parent_socket, recieve_message(socket));
    }
}

void exec(std::istringstream& command_stream, zmq::socket_t& parent_socket, zmq::socket_t& left_socket, 
            zmq::socket_t& right_socket, int& left_pid, int& right_pid, int& id, std::string& request_string, Timer* timer)
{
    std::string subcommand;
    int exec_id;
    command_stream >> exec_id;
    if (exec_id == id)
    {
        command_stream >> subcommand;
        std::string recieve_message = "";
        if (subcommand == "start")
        {
            timer->start_timer();
            recieve_message = "Ok:" + std::to_string(id);
            send_message(parent_socket, recieve_message);
        } 
        else if (subcommand == "stop")
        {
            timer->stop_timer();
            recieve_message = "Ok:" + std::to_string(id);
            send_message(parent_socket, recieve_message);
        }
        else if (subcommand == "time")
        {
            recieve_message = "Ok:" + std::to_string(id) + ": ";
            recieve_message += std::to_string(timer->get_time());
            send_message(parent_socket, recieve_message);
        } 
    } 
    else if (exec_id < id)
    {
        rl_exec(parent_socket, left_socket, exec_id,
                left_pid, request_string);
    } 
    else
    {
        rl_exec(parent_socket, right_socket, exec_id,
                right_pid, request_string);
    }
}


void pingall(zmq::socket_t& parent_socket, int& id, zmq::socket_t& left_socket, zmq::socket_t& right_socket,int& left_pid, int& right_pid)
{
    std::ostringstream res;
    std::string left_res;
    std::string right_res;        
    res << std::to_string(id);
    if (left_pid != 0)
    {
        send_message(left_socket, "pingall");
        left_res = recieve_message(left_socket);
    }
    if (right_pid != 0)
    {
        send_message(right_socket, "pingall");
        right_res = recieve_message(right_socket);
    }
    if (!left_res.empty() && left_res.substr(0, std::min<int>(left_res.size(),5) ) != "Error")
    {
        res << " " << left_res;        
    }
    if ((!right_res.empty()) && (right_res.substr(0, std::min<int>(right_res.size(),5) ) != "Error"))
    {
        res << " "<< right_res;
    }
    send_message(parent_socket, res.str());
}

void kill_children(zmq::socket_t& parent_socket, zmq::socket_t& left_socket, zmq::socket_t& right_socket, int& left_pid, int& right_pid)
{
    if (left_pid == 0 && right_pid == 0)
    {
        send_message(parent_socket, "Ok");
    } 
    else
    {
        if (left_pid != 0)
        {
            send_message(left_socket, "kill_children");
            recieve_message(left_socket);
            kill(left_pid,SIGTERM);
            kill(left_pid,SIGKILL);
        }
        if (right_pid != 0)
        {
            send_message(right_socket, "kill_children");
            recieve_message(right_socket);
            kill(right_pid,SIGTERM);
            kill(right_pid,SIGKILL);
        }
        send_message(parent_socket, "Ok");
    }
}

int main(int argc, char** argv)
{
    Timer timer;
    int id = std::stoi(argv[1]);
    int parent_port = std::stoi(argv[2]);
    zmq::context_t context(3);
    zmq::socket_t parent_socket(context, ZMQ_REP);
    parent_socket.connect(get_port_name(parent_port));
    int left_pid = 0;
    int right_pid = 0;
    int left_id = 0;
    int right_id = 0;
    zmq::socket_t left_socket(context, ZMQ_REQ);
    zmq::socket_t right_socket(context, ZMQ_REQ);
    while(1)
    {
        std::string request_string = recieve_message(parent_socket);
        std::istringstream command_stream(request_string);
        std::string command;
        command_stream >> command;
        if (command == "id")
        {
            std::string parent_string = "Ok:" + std::to_string(id);
            send_message(parent_socket, parent_string);
        } 
        else if (command == "pid")
        {
            std::string parent_string = "Ok:" + std::to_string(getpid());
            send_message(parent_socket, parent_string);
        } 
        else if (command == "create")
        {
            int create_id;
            command_stream >> create_id;
            if (create_id == id)
            {
                std::string message_string = "Error: Already exists";
                send_message(parent_socket, message_string);
            } 
            else if (create_id < id)
            {
                if (left_pid == 0)
                {
                    left_socket.bind(get_port_name(PORT_BASE + create_id));
                    left_pid = fork();
                    rl_create(parent_socket, left_socket, create_id, left_id, left_pid);
                } 
                else
                {
                    send_message(left_socket, request_string);
                    send_message(parent_socket, recieve_message(left_socket));
                }
            } 
            else
            {
                if (right_pid == 0)
                {
                    right_socket.bind(get_port_name(PORT_BASE + create_id));
                    right_pid = fork();
                    rl_create(parent_socket, right_socket, create_id, right_id, right_pid);
                } 
                else
                {
                    send_message(right_socket, request_string);
                    send_message(parent_socket, recieve_message(right_socket));
                }
            }
        } 
        else if (command == "kill")
        {
            int delete_id;
            command_stream >> delete_id;
            if (delete_id < id)
            {
                rl_kill(parent_socket, left_socket, delete_id, left_id, left_pid, request_string);
            } 
            else
            {
                rl_kill(parent_socket, right_socket, delete_id, right_id, right_pid, request_string);
            }
        } 
        else if (command == "exec")
        {
            exec(command_stream, parent_socket, left_socket, right_socket, left_pid, right_pid, id, request_string, &timer);
        }
        else if (command == "pingall")
        {
            pingall(parent_socket, id, left_socket, right_socket, left_pid, right_pid);
        } 
        else if (command == "kill_children")
        {
            kill_children(parent_socket, left_socket, right_socket, left_pid, right_pid); 
        }
        if (parent_port == 0) 
        {
            break;
        }
    }
    return 0;
}

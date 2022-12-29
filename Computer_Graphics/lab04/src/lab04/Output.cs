using System;
using OpenTK;
using OpenTK.Graphics;
using OpenTK.Graphics.OpenGL;
using OpenTK.Input;

namespace lab04
{
    public class Output
    {
        private readonly GameWindow _window;
        private Figure _figure;

        private float _scaling = 10.0f;
        private float _xAngle;
        private float _yAngle;
        private float _lightPositionX = 20.0f;

        public Output(int size)
        {
            _window = new GameWindow(size, size, 
                GraphicsMode.Default, "");

            _window.Load += Window_Load;
            _window.Resize += Window_Resize;
            _window.RenderFrame += Window_RenderFrame;
            _window.UpdateFrame += Window_UpdateFrame;
            _window.KeyDown += Window_KeyDown;
        }

        public void Start()
        {
            _figure = new Figure(2, 20);
            _window.Run(1.0 / 60.0);
        }

        private static void Window_Load(object sender, EventArgs e)
        {
            GL.ClearColor(1.0f, 1.0f, 1.0f, 1.0f);
            GL.Enable(EnableCap.DepthTest);
            GL.Enable(EnableCap.Lighting);
            GL.Enable(EnableCap.Light0);
        }

        private void Window_Resize(object sender, EventArgs e)
        {
            GL.Viewport(0, 0, _window.Width, _window.Height);
            GL.MatrixMode(MatrixMode.Projection);

            GL.LoadIdentity();

            var matrix = 
                Matrix4.CreatePerspectiveFieldOfView((float)Math.PI / 4, 
                    1.0f, 1.0f, 100.0f);
            GL.LoadMatrix(ref matrix);
            GL.MatrixMode(MatrixMode.Modelview);
        }

        private void Window_KeyDown(object sender, KeyboardKeyEventArgs e)
        {
            switch (e.Key)
            {
                case Key.Left:
                    _lightPositionX -= 10.0f;
                    if (_lightPositionX < -360.0f)
                    {
                        _lightPositionX = 360.0f;
                    }
                    break;
                case Key.Right:
                    _lightPositionX += 10.0f;
                    if (_lightPositionX > 360.0f)
                    {
                        _lightPositionX = -360.0f;
                    }
                    break;
                case Key.Up:
                    _figure.Precision++;
                    break;
                case Key.Down:
                    _figure.Precision--;
                    break;
                case Key.Plus:
                    _scaling -= 0.5f;
                    break;
                case Key.Minus:
                    _scaling += 0.5f;
                    break;
                case Key.S:
                    _xAngle += 10.0f;
                    if (_xAngle > 360.0f)
                    {
                        _xAngle = 0.0f;
                    }
                    break;
                case Key.W:
                    _xAngle -= 10.0f;
                    if (_xAngle < 0.0f)
                    {
                        _xAngle = 360.0f;
                    }
                    break;
                case Key.D:
                    _yAngle += 10.0f;
                    if (_yAngle > 360.0f)
                    {
                        _yAngle = 0.0f;
                    }
                    break;
                case Key.A:
                    _yAngle -= 10.0f;
                    if (_yAngle < 0.0f)
                    {
                        _yAngle = 360.0f;
                    }
                    break;
            }
        }

        private void Window_UpdateFrame(object sender, FrameEventArgs e)
        {
            _window.Title = $"©Матвей Попов М8О-308Б-20";
        }

        private void Window_RenderFrame(object sender, FrameEventArgs e)
        {
            GL.LoadIdentity();
            GL.Clear(ClearBufferMask.ColorBufferBit | 
                     ClearBufferMask.DepthBufferBit);

            GL.Translate(0.0, 0.0, -_scaling);

            GL.Rotate(_xAngle, 1.0, 0.0, 0.0);
            GL.Rotate(_yAngle, 0.0, 1.0, 0.0);

            _figure.Draw();

            _figure.LightConfigure(_lightPositionX);

            _window.SwapBuffers();
        }
    }
}

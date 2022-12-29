using System;
using OpenTK.Graphics.OpenGL;

namespace lab04
{
    public class Figure
    {
        private readonly float _radius;
        private int _precision;
        private readonly float _r;
        private readonly float _g;
        private readonly float _b;

        public Figure(float radius, int precision)
        {
            _radius = radius;
            _precision = precision;
            _r = 0.5f;
            _g = 0.5f;
            _b = 0.5f;
        }

        private const int MinPrecision = 3;

        public int Precision
        {
            get => _precision;
            set => _precision = (value < MinPrecision) ? MinPrecision : value;
        }

        public void Draw()
        {
            const float endPhi = (float)Math.PI * 2.0f;
            const float endTheta = (float)Math.PI * 0.5f;
            var dPhi = endPhi / _precision;
            var dTheta = endTheta / _precision;

            for (var pointPhi = 0; pointPhi < _precision; pointPhi++)
            {
                for (var pointTheta = 0; pointTheta < _precision; pointTheta++)
                {
                    var phi = pointPhi * dPhi;
                    var theta = pointTheta * dTheta;
                    var phiT = (pointPhi + 1 == _precision) ? endPhi
                        : (pointPhi + 1) * dPhi;
                    var thetaT = (pointTheta + 1 == _precision) ? endTheta
                        : (pointTheta + 1) * dTheta;

                    float[] p0 = { _radius * (float)Math.Sin(theta) * 
                                   (float)Math.Cos(phi), _radius * 
                        (float)Math.Sin(theta) * (float)Math.Sin(phi), 
                        _radius * (float)Math.Cos(theta) };
                    
                    float[] p1 = { _radius * (float)Math.Sin(thetaT) * 
                                   (float)Math.Cos(phi), _radius * 
                        (float)Math.Sin(thetaT) * (float)Math.Sin(phi), 
                        _radius * (float)Math.Cos(thetaT) };
                    
                    float[] p2 = { _radius * (float)Math.Sin(theta) * 
                                   (float)Math.Cos(phiT), _radius * 
                        (float)Math.Sin(theta) * (float)Math.Sin(phiT), 
                        _radius * (float)Math.Cos(theta) };
                    
                    float[] p3 = { _radius * (float)Math.Sin(thetaT) * 
                                   (float)Math.Cos(phiT), _radius * 
                        (float)Math.Sin(thetaT) * (float)Math.Sin(phiT), 
                        _radius * (float)Math.Cos(thetaT) };

                    GL.Begin(PrimitiveType.Triangles);
                    GL.Normal3(p0[0] / _radius, p0[1] / _radius, 
                        p0[2] / _radius);
                    
                    GL.Vertex3(p0[0], p0[1], p0[2]);
                    
                    GL.Normal3(p2[0] / _radius, p2[1] / _radius, 
                        p2[2] / _radius);
                    
                    GL.Vertex3(p2[0], p2[1], p2[2]);
                    
                    GL.Normal3(p1[0] / _radius, p1[1] / _radius, 
                        p1[2] / _radius);
                    
                    GL.Vertex3(p1[0], p1[1], p1[2]);
                    
                    GL.Normal3(p3[0] / _radius, p3[1] / _radius, 
                        p3[2] / _radius);
                    
                    GL.Vertex3(p3[0], p3[1], p3[2]);
                    
                    GL.Normal3(p1[0] / _radius, p1[1] / _radius, 
                        p1[2] / _radius);
                    
                    GL.Vertex3(p1[0], p1[1], p1[2]);
                    
                    GL.Normal3(p2[0] / _radius, p2[1] / _radius, 
                        p2[2] / _radius);
                    
                    GL.Vertex3(p2[0], p2[1], p2[2]);
                    
                    GL.Normal3(p0[0] / _radius, p0[1] / _radius, 0);
                    
                    GL.Vertex3(p0[0], p0[1], 0);
                    
                    GL.Normal3(p2[0] / _radius, p2[1] / _radius, 0);
                    
                    GL.Vertex3(p2[0], p2[1], 0);
                    
                    GL.Normal3(p1[0] / _radius, p1[1] / _radius, 0);
                    
                    GL.Vertex3(p1[0], p1[1], 0);
                    
                    GL.Normal3(p3[0] / _radius, p3[1] / _radius, 0);
                    
                    GL.Vertex3(p3[0], p3[1], 0);
                    
                    GL.Normal3(p1[0] / _radius, p1[1] / _radius, 0);
                    
                    GL.Vertex3(p1[0], p1[1], 0);
                    
                    GL.Normal3(p2[0] / _radius, p2[1] / _radius, 0);
                    
                    GL.Vertex3(p2[0], p2[1], 0);
                    
                    GL.End();
                }
            }
        }

        public void LightConfigure(float lpx)
        {
            float[] lightPosition = {lpx, 20, 80};
            float[] lightDiffuse = {_r, _g, _b};

            GL.Light(LightName.Light0, LightParameter.Position, lightPosition);
            GL.Light(LightName.Light0, LightParameter.Diffuse, lightDiffuse);
            GL.Light(LightName.Light0, LightParameter.Ambient, lightDiffuse);
        }
    }
}

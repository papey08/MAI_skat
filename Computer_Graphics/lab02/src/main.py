import tkinter as tk
import matrix as mtx
import camera as cm
import roberts as rb
import object as obj
from functools import partial

def MVPMatrix(object):
    return mtx.ModelMatrix(object) @ camera.ViewMatrix() @ camera.projectionMatrix()

def OMVPMatrix(object):
    return mtx.ModelMatrix(object) @ camera.ViewMatrix() @ camera.orthogonalMatrix()

def redraw(event, object):
    camera.height = canvas.winfo_height()
    camera.width = canvas.winfo_width()
    camera.v_fov = camera.h_fov * (camera.height/camera.width)
    drawOrthogonalProjection(object)

def changeStartModel(event, object):
    object.scretchStartModel(sx=scaleX.get(), sy=scaleY.get(), sz=scaleZ.get())
    drawOrthogonalProjection(object)

def changeScaleModel(event, object):
    object.scale = scale.get()
    drawOrthogonalProjection(object)

def keyMove(event, object):
    if event.keycode == 37:
        object.rotY += 3
    elif event.keycode == 38:
        object.rotX -= 3
    elif event.keycode == 39:
        object.rotY -= 3
    elif event.keycode == 40:
        object.rotX += 3
    drawOrthogonalProjection(object)

def drawLine(a, b):
    canvas.create_line(a[0], a[1], b[0], b[1], width=2)

def drawOrthogonalProjection(object):
    canvas.delete("all")
    robAlgo = rb.RobertsAlgo(object.points @ mtx.ModelMatrix(object), object.polygon, camera.ViewMatrix(), camera.position)
    vertex = object.points @ OMVPMatrix(object)
    for i in range(len(object.polygon)):
        if robAlgo[i] <= 0:
            continue
        for j in range(len(object.polygon[i])):
            drawLine(vertex[object.polygon[i][j]], vertex[object.polygon[i][(j + 1) % len(object.polygon[i])]])

def drawPerspectiveProjection(object):
    vertex = object.points @ MVPMatrix()
    vertex /= vertex[:, -1].reshape(-1, 1)
    vertex = vertex @ camera.toScreenMatrix()

    for i in range(len(polygon)):
        for j in range(len(polygon[i])):
            a = vertex[object.polygon[i][j]]
            b = vertex[object.polygon[i][(j+1)%len(object.polygon[i])]]
            if not(a and b):
                continue
            drawLine(a, b)

if __name__=="__main__":
    window = tk.Tk()
    window.title("CG_lab02_var22")
    window.columnconfigure(0, weight=4, minsize=550)
    window.columnconfigure([1, 2], weight=1, minsize=10)
    window.rowconfigure([0, 1, 2, 3], weight=1, minsize=100)
    canvas = tk.Canvas(window, bg='white')
    frame = tk.Frame(window, relief="sunken", borderwidth=3)
    labelscale = tk.Label(master=frame, text='Scale')
    scale = tk.Scale(frame, orient='horizontal', resolution=1, from_=1, to=10)
    canvas.grid(row=0, column=0, sticky="nsew", rowspan=4)
    frame.grid(row=0, column=1, sticky="nsew", rowspan=4, columnspan=2)
    scale.grid(row=3, column=2, sticky="nsew", columnspan=3)
    window.geometry("+550+150")
    window.minsize(window.winfo_width(), window.winfo_height())
    camera = cm.Camera(canvas)
    trapeze = obj.Object(sx=10, sy=10, sz=10)
    drawOrthogonalProjection(trapeze)
    scale.bind("<ButtonRelease-1>", partial(changeScaleModel, object=trapeze))
    window.bind("<Key>", partial(keyMove, object=trapeze))
    window.bind("<Configure>", partial(redraw, object=trapeze))
    window.mainloop()

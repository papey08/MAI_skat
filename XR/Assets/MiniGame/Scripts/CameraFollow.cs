using UnityEngine;

[RequireComponent(typeof(Camera))]
public class CameraFollow : MonoBehaviour
{
    public Transform target;
    public float followSmoothSpeed;

    public float baseZoom;
    public float scaleZoomFactor;
    public float zoomSmoothSpeed;

    private Camera cam;
    private float targetZoom;

    private void Start()
    {
        cam = GetComponent<Camera>();
        targetZoom = baseZoom;
        cam.orthographicSize = baseZoom;
    }

    private void LateUpdate()
    {
        if (target == null) return;

        float scaleX = target.localScale.x;

        Vector3 targetPosition = target.position;
        targetPosition.z = -scaleX * 10f;

        transform.position = Vector3.Lerp(transform.position, targetPosition, followSmoothSpeed);

        targetZoom = baseZoom + scaleX * scaleZoomFactor;
        cam.orthographicSize = Mathf.Lerp(cam.orthographicSize, targetZoom, Time.deltaTime * zoomSmoothSpeed);
    }
}

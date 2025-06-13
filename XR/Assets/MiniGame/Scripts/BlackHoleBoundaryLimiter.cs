using UnityEngine;

[RequireComponent(typeof(BlackHoleController))]
public class BlackHoleBoundaryLimiter : MonoBehaviour
{
    private void LateUpdate()
    {
        if (BoundaryManager.Instance == null) return;

        float currentScale = transform.localScale.x;
        float maxRadius = BoundaryManager.Instance.GetBoundaryRadius(currentScale);

        Vector2 center = Vector2.zero;
        Vector2 currentPosition = transform.position;
        Vector2 direction = currentPosition - center;

        if (direction.magnitude > maxRadius)
        {
            transform.position = center + direction.normalized * maxRadius;
        }
    }
}

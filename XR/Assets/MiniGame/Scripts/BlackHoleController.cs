using UnityEngine;

public class BlackHoleController : MonoBehaviour
{
    public float distancePerSecondFactor;

    private void Update()
    {
        HandleMovement();
    }

    private void HandleMovement()
    {
        float moveX = Input.GetAxisRaw("Horizontal");
        float moveY = Input.GetAxisRaw("Vertical");

        Vector3 inputDirection = new Vector3(moveX, moveY, 0f).normalized;

        float scaleX = transform.localScale.x;
        float moveSpeed = scaleX * distancePerSecondFactor;

        transform.position += inputDirection * moveSpeed * Time.deltaTime;
    }

    public void Absorb(float absorbedScaleX)
    {
        float currentScale = transform.localScale.x;

        float currentArea = currentScale * currentScale;
        float absorbedArea = absorbedScaleX * absorbedScaleX / 4f;

        float newArea = currentArea + absorbedArea;
        float newScale = Mathf.Sqrt(newArea);

        transform.localScale = new Vector3(newScale, newScale, 1f);
    }

    public float GetVisualRadius()
    {
        return transform.localScale.x / 2f;
    }
}

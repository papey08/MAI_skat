using UnityEngine;

public class DisableAfterDelay : MonoBehaviour
{
    public float delay;

    void Start()
    {
        Invoke(nameof(DisableObject), delay);
    }

    void DisableObject()
    {
        gameObject.SetActive(false);
    }
}

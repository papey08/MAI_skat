using UnityEngine;

public class BossDeathNotifier : MonoBehaviour
{
    public static event System.Action OnBossDeath;

    private void OnDestroy()
    {
        OnBossDeath?.Invoke();
    }
}

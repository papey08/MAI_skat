using UnityEngine;
using System;

public class EnemyDeathNotifier : MonoBehaviour
{
    public event Action OnDeath;

    private void OnDestroy()
    {
        OnDeath?.Invoke();
    }
}

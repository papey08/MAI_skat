using UnityEngine;
using System.Collections.Generic;

public class BoundaryManager : MonoBehaviour
{
    public static BoundaryManager Instance { get; private set; }

    public List<ScaleRadiusPair> boundaries = new List<ScaleRadiusPair>();

    private Dictionary<float, float> boundaryDict = new Dictionary<float, float>();

    [System.Serializable]
    public struct ScaleRadiusPair
    {
        public float scaleX;
        public float radius;
    }

    private void Awake()
    {
        if (Instance != null && Instance != this)
        {
            Destroy(gameObject);
            return;
        }

        Instance = this;

        foreach (var pair in boundaries)
        {
            if (!boundaryDict.ContainsKey(pair.scaleX))
            {
                boundaryDict.Add(pair.scaleX, pair.radius);
            }
        }
    }

    public float GetBoundaryRadius(float currentScaleX)
    {
        float closestKey = -1f;
        float minDiff = float.MaxValue;

        foreach (float key in boundaryDict.Keys)
        {
            float diff = Mathf.Abs(key - currentScaleX);
            if (diff < minDiff)
            {
                minDiff = diff;
                closestKey = key;
            }
        }

        if (closestKey != -1f)
            return boundaryDict[closestKey];
        return 20f;
    }
}

using UnityEngine;

public class LandingZoneActivator : MonoBehaviour
{
    public GameObject landingZone;
    public GameObject landingInstructions;

    private void OnEnable()
    {
        BossDeathNotifier.OnBossDeath += ActivateLandingZone;
    }

    private void OnDisable()
    {
        BossDeathNotifier.OnBossDeath -= ActivateLandingZone;
    }

    private void ActivateLandingZone()
    {
        if (landingZone != null)
        {
            landingZone.SetActive(true);
        }

        if (landingInstructions != null)
        {
            landingInstructions.SetActive(true);
        }
    }
}

using UnityEngine;

public class BossMonitor : MonoBehaviour
{
    public GameObject boss;

    public GameObject landingZone;
    public GameObject instruction;

    private bool landingZoneActivated = false;
    private bool instructionActivated = false;

    void Update()
    {
        if (!landingZoneActivated && boss == null)
        {
            ActivateLandingZone();
        }
        if (!instructionActivated && boss != null)
        {
            ActivateInstruction();
        }
    }

    void ActivateLandingZone()
    {
        if (landingZone != null)
        {
            landingZone.SetActive(true);
            landingZoneActivated = true;
        }
    }

    void ActivateInstruction()
    {
        if (instruction != null)
        {
            instruction.SetActive(true);
            instructionActivated = true;
        }
    }
}

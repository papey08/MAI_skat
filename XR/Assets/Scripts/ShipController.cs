using UnityEngine;
using UnityEngine.UI;

public class ShipController : MonoBehaviour
{
    [Header("Movement")]
    public float rotationSpeed;
    public float thrustForce;
    public float maxSpeed;
    public float zStabilizationSpeed;
    public float collisionPushBackForce;

    [Header("Fuel")]
    public float maxFuel;
    public float currentFuel;
    public float fuelConsumptionRate;
    public float refuelRate;
    public Slider fuelSlider;

    [Header("Refuel")]
    public RefuelZone refuelZone;
    public GameObject refuelIndicator;

    [Header("RedLaser")]
    public GameObject redLaserPrefab;
    public Transform redLaserPoint;
    public float laserFuelCost;

    [Header("BlueLaser")]
    public float maxBlueLaser;
    public float currentBlueLaser;
    public Slider blueLaserSlider;

    public GameObject blueLaserPrefab;
    public Transform blueLaserPoint;
    public float blueLaserCost;

    [Header("YellowLaser")]
    public float maxYellowLaser;
    public float currentYellowLaser;
    public Slider yellowLaserSlider;

    public GameObject yellowLaserPrefab;
    public Transform yellowLaserPoint;
    public float yellowLaserCost;

    [Header("Pause")]
    public GameObject pauseSign;

    private Rigidbody rb;
    private bool isPaused;



    private KeyCode[] konamiCode = new KeyCode[] {
        KeyCode.UpArrow, KeyCode.UpArrow,
        KeyCode.DownArrow, KeyCode.DownArrow,
        KeyCode.LeftArrow, KeyCode.RightArrow,
        KeyCode.LeftArrow, KeyCode.RightArrow,
        KeyCode.B, KeyCode.A
    };
    private int konamiIndex = 0;
    private bool cheatMode = false;

    void Start()
    {
        rb = GetComponent<Rigidbody>();
        rb.useGravity = false;
        rb.angularDamping = 5f;
        rb.linearDamping = 1f;
        rb.maxAngularVelocity = 2f;

        currentFuel = maxFuel;
        UpdateFuelUI();
        UpdateBlueLaserUI();

        if (refuelIndicator != null)
            refuelIndicator.SetActive(false);

        if (pauseSign != null)
            pauseSign.SetActive(false);
    }

    void Update()
    {
        HandleRedLaserShooting();
        HandleBlueLaserShooting();
        HandleYellowLaserShooting();
        HandleRotation();
        HandleThrust();
        LimitSpeed();
        StabilizeZRotation();
        HandleRefueling();
        UpdateFuelUI();
        HandlePause();
        HandleKonamiCode();
    }

    void HandlePause()
    {
        if (Input.GetKeyDown(KeyCode.P))
        {
            TogglePause();
            return;
        }
    }

    void TogglePause()
    {
        isPaused = !isPaused;
        Time.timeScale = isPaused ? 0f : 1f;
        if (rb != null)
            rb.isKinematic = isPaused;
        

        if (pauseSign != null)
            pauseSign.SetActive(isPaused);
    }

    void HandleRotation()
    {
        float pitch = 0f;
        float yaw = 0f;

        if (Input.GetKey(KeyCode.W)) pitch = 1f;
        if (Input.GetKey(KeyCode.S)) pitch = -1f;
        if (Input.GetKey(KeyCode.A)) yaw = -1f;
        if (Input.GetKey(KeyCode.D)) yaw = 1f;

        Vector3 rotation = new Vector3(pitch, yaw, 0f) * rotationSpeed * Time.deltaTime;
        Quaternion deltaRotation = Quaternion.Euler(rotation);
        rb.MoveRotation(rb.rotation * deltaRotation);
    }

    void HandleThrust()
    {
        bool forward = Input.GetKey(KeyCode.Space);
        bool backward = Input.GetKey(KeyCode.LeftShift);

        if ((forward || backward) && currentFuel > 0f)
        {
            Vector3 direction = forward ? transform.forward : -transform.forward;
            rb.AddForce(direction * thrustForce, ForceMode.Acceleration);

            if (!cheatMode)
            {
                currentFuel -= fuelConsumptionRate * Time.deltaTime;
                currentFuel = Mathf.Max(currentFuel, 0f);
            }
        }
    }

    void HandleRefueling()
    {
        if (refuelZone == null) return;

        bool inRange = refuelZone.IsPlayerInRange(transform);

        if (refuelIndicator != null)
            refuelIndicator.SetActive(inRange);

        if (inRange && Input.GetKey(KeyCode.F) && currentFuel < maxFuel && !cheatMode)
        {
            currentFuel += refuelRate * Time.deltaTime;
            currentFuel = Mathf.Min(currentFuel, maxFuel);
        }
    }

    void LimitSpeed()
    {
        if (rb.linearVelocity.magnitude > maxSpeed)
            rb.linearVelocity = rb.linearVelocity.normalized * maxSpeed;
    }

    void StabilizeZRotation()
    {
        Vector3 currentEuler = rb.rotation.eulerAngles;
        float z = currentEuler.z;
        if (z > 180f) z -= 360f;
        z = Mathf.Lerp(z, 0f, zStabilizationSpeed * Time.deltaTime);
        Quaternion targetRotation = Quaternion.Euler(currentEuler.x, currentEuler.y, z);
        rb.MoveRotation(Quaternion.Slerp(rb.rotation, targetRotation, zStabilizationSpeed * Time.deltaTime));
    }

    void OnCollisionEnter(Collision collision)
    {
        rb.angularVelocity = Vector3.zero;
        rb.linearVelocity = Vector3.zero;
        rb.AddForce(-transform.forward * collisionPushBackForce, ForceMode.VelocityChange);

        if (collision.gameObject.CompareTag("Enemy"))
        {
            Destroy(gameObject);
        }
    }

    void UpdateFuelUI()
    {
        if (fuelSlider != null)
            fuelSlider.value = currentFuel / maxFuel;
    }

    void UpdateBlueLaserUI()
    {
        if (blueLaserSlider != null)
        {
            blueLaserSlider.maxValue = maxBlueLaser;
            blueLaserSlider.value = currentBlueLaser;
        }
    }

    void UpdateYellowLaserUI()
    {
        if (yellowLaserSlider != null)
        {
            yellowLaserSlider.maxValue = maxYellowLaser;
            yellowLaserSlider.value = currentYellowLaser;
        }
    }

    void HandleRedLaserShooting()
    {
        if (Input.GetKeyDown(KeyCode.Q) && currentFuel >= laserFuelCost)
        {
            Instantiate(redLaserPrefab, redLaserPoint.position, redLaserPoint.rotation);
            if (!cheatMode)
            {
                currentFuel -= laserFuelCost;
                UpdateFuelUI();
            }
        }
    }

    void HandleBlueLaserShooting()
    {
        if (Input.GetKeyDown(KeyCode.E) && currentBlueLaser >= blueLaserCost)
        {
            Instantiate(blueLaserPrefab, blueLaserPoint.position, blueLaserPoint.rotation);
            if (!cheatMode)
            {
                currentBlueLaser -= blueLaserCost;
                UpdateBlueLaserUI();
            }
        }
    }

    void HandleYellowLaserShooting()
    {
        if (Input.GetKeyDown(KeyCode.R) && currentYellowLaser >= yellowLaserCost)
        {
            Instantiate(yellowLaserPrefab, yellowLaserPoint.position, yellowLaserPoint.rotation);
            if (!cheatMode)
            {
                currentYellowLaser -= yellowLaserCost;
                UpdateYellowLaserUI();
            }
        }
    }

    public void AddBlueLaserCharge()
    {
        if (currentBlueLaser < maxBlueLaser && !cheatMode)
        {
            currentBlueLaser++;
            UpdateBlueLaserUI();
        }
    }

    public void AddYellowLaserCharge()
    {
        if (currentYellowLaser < maxYellowLaser && !cheatMode)
        {
            currentYellowLaser++;
            UpdateYellowLaserUI();
        }
    }

    void HandleKonamiCode()
    {
        if (konamiIndex < konamiCode.Length && Input.GetKeyDown(konamiCode[konamiIndex]))
        {
            konamiIndex++;
            if (konamiIndex >= konamiCode.Length)
            {
                cheatMode = true;
                currentFuel = maxFuel;
                currentBlueLaser = maxBlueLaser;
                currentYellowLaser = maxYellowLaser;
                if (fuelSlider != null) fuelSlider.value = 1f;
                if (blueLaserSlider != null) blueLaserSlider.value = maxBlueLaser;
                if (yellowLaserSlider != null) yellowLaserSlider.value = maxYellowLaser;
            }
        }
        else if (Input.anyKeyDown && !Input.GetKeyDown(konamiCode[konamiIndex]))
        {
            konamiIndex = 0;
        }
    }
}

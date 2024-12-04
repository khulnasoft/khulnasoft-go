package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/khulnasoft/khulnasoft-go"
	"github.com/urfave/cli/v2"
)

func formatDNSRecord(record khulnasoft.DNSRecord) []string {
	return []string{
		record.ID,
		record.Name,
		record.Type,
		record.Content,
		strconv.FormatInt(int64(record.TTL), 10),
		strconv.FormatBool(record.Proxiable),
		strconv.FormatBool(*record.Proxied),
	}
}

func dnsCreate(c *cli.Context) error {
	if err := checkFlags(c, "zone", "name", "type", "content"); err != nil {
		return err
	}
	zone := c.String("zone")
	name := c.String("name")
	rtype := c.String("type")
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")
	priority := uint16(c.Uint("priority"))

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	record := khulnasoft.CreateDNSRecordParams{
		Name:     name,
		Type:     strings.ToUpper(rtype),
		Content:  content,
		TTL:      ttl,
		Proxied:  &proxy,
		Priority: &priority,
	}
	result, err := api.CreateDNSRecord(context.Background(), khulnasoft.ZoneIdentifier(zoneID), record)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating DNS record: ", err)
		return err
	}

	output := [][]string{
		formatDNSRecord(result),
	}

	writeTable(c, output, "ID", "Name", "Type", "Content", "TTL", "Proxiable", "Proxy")

	return nil
}

func dnsCreateOrUpdate(c *cli.Context) error {
	if err := checkFlags(c, "zone", "name", "type", "content"); err != nil {
		fmt.Println(err)
		return err
	}
	zone := c.String("zone")
	name := c.String("name")
	rtype := strings.ToUpper(c.String("type"))
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")
	priority := uint16(c.Uint("priority"))

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error updating DNS record: ", err)
		return err
	}

	records, _, err := api.ListDNSRecords(context.Background(), khulnasoft.ZoneIdentifier(zoneID), khulnasoft.ListDNSRecordsParams{Name: name + "." + zone})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error fetching DNS records: ", err)
		return err
	}

	var result khulnasoft.DNSRecord
	if len(records) > 0 {
		// Record exists - find the ID and update it.
		// This is imprecise without knowing the original content; if a label
		// has multiple RRs we'll just update the first one.
		for _, r := range records {
			if r.Type == rtype {
				rr := khulnasoft.UpdateDNSRecordParams{}
				rr.ID = r.ID
				rr.Type = r.Type
				rr.Content = content
				rr.TTL = ttl
				rr.Proxied = &proxy
				rr.Priority = &priority

				result, err = api.UpdateDNSRecord(context.Background(), khulnasoft.ZoneIdentifier(zoneID), rr)
				if err != nil {
					fmt.Println("Error updating DNS record:", err)
					return err
				}
			}
		}
	} else {
		// Record doesn't exist - create it
		rr := khulnasoft.CreateDNSRecordParams{
			Name:     name,
			Type:     rtype,
			Content:  content,
			TTL:      ttl,
			Proxied:  &proxy,
			Priority: &priority,
		}

		// TODO: Print the response.
		result, err = api.CreateDNSRecord(context.Background(), khulnasoft.ZoneIdentifier(zoneID), rr)
		if err != nil {
			fmt.Println("Error creating DNS record:", err)
			return err
		}
	}

	output := [][]string{
		formatDNSRecord(result),
	}

	writeTable(c, output, "ID", "Name", "Type", "Content", "TTL", "Proxiable", "Proxy")

	return nil
}

func dnsUpdate(c *cli.Context) error {
	if err := checkFlags(c, "zone", "id"); err != nil {
		fmt.Println(err)
		return err
	}
	zone := c.String("zone")
	recordID := c.String("id")
	name := c.String("name")
	rtype := c.String("type")
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")
	priority := uint16(c.Uint("priority"))

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	record := khulnasoft.UpdateDNSRecordParams{
		ID:       recordID,
		Name:     name,
		Type:     strings.ToUpper(rtype),
		Content:  content,
		TTL:      ttl,
		Proxied:  &proxy,
		Priority: &priority,
	}
	_, err = api.UpdateDNSRecord(context.Background(), khulnasoft.ZoneIdentifier(zoneID), record)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error updating DNS record: ", err)
		return err
	}

	return nil
}

func dnsDelete(c *cli.Context) error {
	if err := checkFlags(c, "zone", "id"); err != nil {
		fmt.Println(err)
		return err
	}
	zone := c.String("zone")
	recordID := c.String("id")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = api.DeleteDNSRecord(context.Background(), khulnasoft.ZoneIdentifier(zoneID), recordID)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error deleting DNS record: ", err)
		return err
	}

	return nil
}
